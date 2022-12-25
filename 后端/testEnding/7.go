package main

import (
	"log"
)

var warehouseList []Warehouse



// 仓库结构体
type Warehouse struct {
	Name  string  // 仓库名
	Goods []Goods // 仓库货物
}

// 货物结构体
type Goods struct {
	Name     string // 货物名
	Quantity int    // 数量
	Status   bool   // 状态
}

// 新增仓库
func AddWarehouse(name string, goods []Goods) {
	// 创建仓库
	warehouse := Warehouse{
		Name:  name,
		Goods: goods,
	}
	// 将仓库添加到仓库列表中
	warehouseList = append(warehouseList, warehouse)
}

// 上架货物
func PutOnGoods(warehouseName string, goodsName string, quantity int) {
	// 遍历仓库列表，找到目标仓库
	for _, warehouse := range warehouseList {
		if warehouse.Name == warehouseName {
			// 遍历仓库货物，找到目标货物
			for i, good := range warehouse.Goods {
				if good.Name == goodsName {
					// 设置货物状态为上架
					warehouse.Goods[i].Status = true
					// 设置货物数量
					warehouse.Goods[i].Quantity = quantity
					break
				}
			}
			break
		}
	}
}

// 下架货物
func PutOffGoods(warehouseName string, goodsName string) {
	// 遍历仓库列表，找到目标仓库
	for _, warehouse := range warehouseList {
		if warehouse.Name == warehouseName {
			// 遍历仓库货物，找到目标货物
			for i, good := range warehouse.Goods {
				if good.Name == goodsName {
					// 设置货物状态为下架
					warehouse.Goods[i].Status = false
					break
				}
			}
			break
		}
	}
}

// 入库货物
func InGoods(warehouseName string, goodsName string, quantity int) {
	// 遍历仓库列表，找到目标仓库
	for _, warehouse := range warehouseList {
		if warehouse.Name == warehouseName {
			// 遍历仓库货物，找到目标货物
			//设置flag
			flag := 0
			for i, good := range warehouse.Goods {
				if good.Name == goodsName {
					// 设置货物数量
					warehouse.Goods[i].Quantity += quantity
					flag = 1
					break
				}
			}
			//未找到商品,添加商品
			if flag == 0 {
				var good Goods
				good.Status = false
				good.Name = goodsName
				good.Quantity = quantity
				warehouse.Goods = append(warehouse.Goods, good)
			}
			break
		}
	}
}

// 出库货物
func OutGoods(warehouseName string, goodsName string, quantity int) {
	// 遍历仓库列表，找到目标仓库
	for _, warehouse := range warehouseList {
		if warehouse.Name == warehouseName {
			// 遍历仓库货物，找到目标货物
			//设置flag
			flag := 0
			for i, good := range warehouse.Goods {
				if good.Name == goodsName {
					// 设置货物数量
					warehouse.Goods[i].Quantity -= quantity
					flag = 1
					break
				}
			}
			//未找到商品
			if flag == 0 {
				log.Println("该商品不存在")
			}
			break
		}
	}
}

