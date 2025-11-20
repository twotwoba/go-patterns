package visitor

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestStreamVisitor 测试 StreamVisitor 是否能正确解析并访问 Info 对象
func TestStreamVisitor(t *testing.T) {
	// 准备包含 JSON 对象的 Reader
	jsonStream := `
		{"Name": "resource-1", "Namespace": "ns-1"}
		{"Name": "resource-2", "Namespace": "ns-2"}
	`
	reader := strings.NewReader(jsonStream)

	// 创建 StreamVisitor
	sv := &StreamVisitor{Reader: reader, Source: "test-stream"}

	// 准备一个 VisitorFunc 来收集访问到的 Info
	var visitedInfos []*Info
	visitorFunc := func(info *Info, err error) error {
		if err != nil {
			return err
		}
		visitedInfos = append(visitedInfos, info)
		return nil
	}

	// 执行 Visit
	if err := sv.Visit(visitorFunc); err != nil {
		t.Fatalf("StreamVisitor.Visit() returned an unexpected error: %v", err)
	}

	// 验证结果
	if len(visitedInfos) != 2 {
		t.Fatalf("expected to visit 2 infos, but got %d", len(visitedInfos))
	}
	if visitedInfos[0].Name != "resource-1" || visitedInfos[0].Namespace != "ns-1" {
		t.Errorf("unexpected info[0]: got %+v", visitedInfos[0])
	}
	if visitedInfos[1].Name != "resource-2" || visitedInfos[1].Namespace != "ns-2" {
		t.Errorf("unexpected info[1]: got %+v", visitedInfos[1])
	}
}

// TestURLVisitor 测试 URLVisitor 是否能从 HTTP 端点获取并访问 Info
func TestURLVisitor(t *testing.T) {
	// 创建一个 httptest 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Name": "resource-from-url", "Namespace": "ns-url"}`)
	}))
	defer server.Close()

	serverURL, _ := url.Parse(server.URL)

	// 创建 URLVisitor
	uv := &URLVisitor{URL: serverURL}

	// 准备 VisitorFunc
	var visitedInfo *Info
	visitorFunc := func(info *Info, err error) error {
		if err != nil {
			return err
		}
		visitedInfo = info
		return nil
	}

	// 执行 Visit
	if err := uv.Visit(visitorFunc); err != nil {
		t.Fatalf("URLVisitor.Visit() returned an unexpected error: %v", err)
	}

	// 验证结果
	if visitedInfo == nil {
		t.Fatal("URLVisitor did not visit any info")
	}
	if visitedInfo.Name != "resource-from-url" || visitedInfo.Namespace != "ns-url" {
		t.Errorf("unexpected info from URL: got %+v", visitedInfo)
	}
}

// mockVisitor 用于测试 EagerVisitorList
type mockVisitor struct {
	infosToVisit []*Info
	visitErr     error
	funcErr      error
}

// Visit 实现了 Visitor 接口
func (m *mockVisitor) Visit(fn VisitorFunc) error {
	if m.visitErr != nil {
		return m.visitErr
	}
	for _, info := range m.infosToVisit {
		if err := fn(info, m.funcErr); err != nil {
			return err
		}
	}
	return nil
}

// TestEagerVisitorList 测试聚合访问者列表
func TestEagerVisitorList(t *testing.T) {
	// 创建两个 mock visitor
	visitor1 := &mockVisitor{
		infosToVisit: []*Info{
			{Name: "v1-res1", Namespace: "ns1"},
			{Name: "v1-res2", Namespace: "ns1"},
		},
	}
	visitor2 := &mockVisitor{
		infosToVisit: []*Info{
			{Name: "v2-res1", Namespace: "ns2"},
		},
	}
	// 创建一个会返回错误的 visitor
	visitor3_with_err := &mockVisitor{
		visitErr: fmt.Errorf("visitor3 failed"),
	}

	// 测试用例
	t.Run("successful visit", func(t *testing.T) {
		list := EagerVisitorList{visitor1, visitor2}
		var count int
		visitorFunc := func(info *Info, err error) error {
			count++
			return nil
		}

		if err := list.Visit(visitorFunc); err != nil {
			t.Fatalf("EagerVisitorList.Visit() failed: %v", err)
		}
		if count != 3 {
			t.Errorf("expected to visit 3 items, but counter is %d", count)
		}
	})

	t.Run("aggregate errors", func(t *testing.T) {
		list := EagerVisitorList{visitor1, visitor3_with_err}
		visitorFunc := func(info *Info, err error) error {
			// 让第一次调用返回错误
			if info.Name == "v1-res1" {
				return fmt.Errorf("func failed on %s", info.Name)
			}
			return nil
		}

		err := list.Visit(visitorFunc)
		if err == nil {
			t.Fatal("expected an error from EagerVisitorList.Visit(), but got nil")
		}

		errStr := err.Error()
		if !strings.Contains(errStr, "func failed on v1-res1") {
			t.Errorf("error message does not contain function error")
		}
		if !strings.Contains(errStr, "visitor3 failed") {
			t.Errorf("error message does not contain visitor error")
		}
	})
}
