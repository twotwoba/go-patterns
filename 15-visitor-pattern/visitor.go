package visitor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
允许一个或者多个操作应用到对象上，解耦操作和对象本身

	表面：某个对象执行了一个方法
	内部：对象内部调用了多个方法，最后统一返回结果

设计思路：
 1. 对象Visitor interface
 2. Vistor对应的操作VisitorFunc
 3. 封装多个vistor []Visitor为统一的一个
*/
type Info struct {
	Namespace string
	Name      string
}

// visitor 接口
type Visitor interface {
	Visit(VisitorFunc) error
}

// VisitorFunc对应这个对象的方法，也就是定义中的“操作”
type VisitorFunc func(*Info, error) error

// 将多个[]Visitor封装为一个Visitor
type EagerVisitorList []Visitor

// 返回的错误暂存到[]error中，统一聚合
func (l EagerVisitorList) Visit(fn VisitorFunc) error {
	var errs []string
	for i := range l {
		err := l[i].Visit(func(info *Info, err error) error {
			if err != nil {
				errs = append(errs, err.Error())
				return nil
			}
			if err := fn(info, nil); err != nil {
				errs = append(errs, err.Error())
			}
			return nil
		})
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("errors visiting list: \n%s", strings.Join(errs, "\n"))
	}
	return nil
}

type StreamVisitor struct {
	io.Reader
	Source string
}

// 实现 Visit，从 Reader 中解析 Info 对象并调用 fn
func (s *StreamVisitor) Visit(fn VisitorFunc) error {
	decoder := json.NewDecoder(s.Reader)
	for {
		var info Info
		if err := decoder.Decode(&info); err != nil {
			if err == io.EOF {
				break
			}
			return fn(nil, fmt.Errorf("error decoding from source %q: %w", s.Source, err))
		}
		if err := fn(&info, nil); err != nil {
			return err
		}
	}
	return nil
}

// url visit
type URLVisitor struct {
	URL              *url.URL
	HttpAttemptCount int
}

// 实现 Visit, 获取 URL 内容并委托给 StreamVisitor
func (u *URLVisitor) Visit(fn VisitorFunc) error {
	resp, err := http.Get(u.URL.String())
	if err != nil {
		return fn(nil, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fn(nil, fmt.Errorf("bad status code %d from %s", resp.StatusCode, u.URL))
	}

	// 委托给 StreamVisitor 来处理响应体
	sv := &StreamVisitor{
		Reader: resp.Body,
		Source: u.URL.String(),
	}
	return sv.Visit(fn)
}
