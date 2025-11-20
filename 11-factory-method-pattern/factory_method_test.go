package abstractfactory

import "testing"

func TestConCreteFactory_CreateProduct(t *testing.T) {
	conFactory := &ConCreteFactory1{}
	product := conFactory.CreateProduct1()
	conProduct := product.(*ConcreteProduct)

	if conProduct.Name != "KG" {
		t.Error("abstract factory can not create the concreate product")
	}

	conFactory2 := &ConCreteFactory2{}
	product2 := conFactory2.CreateProduct2()
	conProduct2 := product2.(*ConcreteProduct)
	if conProduct2.Name != "KG2" {
		t.Error("abstract factory can not create the concreate product")
	}

}
