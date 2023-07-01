package main

import (
	"fmt"

	"github.com/0x5ab/gomj/gameplay"
	"github.com/0x5ab/gomj/ruleset"
	"github.com/0x5ab/gomj/ruleset/ruleset_jp"
	"github.com/0x5ab/gomj/tiles"
)

func main() {
	t, err := tiles.ParseTiles("1z1z1z2z2z2z5z")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)
	hand := gameplay.Hand{Tiles: t, Fulu: []gameplay.Fulu{{StartTile: tiles.Bai, Type: gameplay.Peng}, {StartTile: tiles.Fa, Type: gameplay.Peng}}}
	fmt.Printf("%s\n", hand.String())
	fmt.Println("能胡: ")
	for _, tt := range tiles.AllTiles {
		huways := ruleset.CanHu(&ruleset_jp.YiZhongRuleset, hand, gameplay.PlayedTile{Tile: tt})
		for _, huway := range huways.Ways {
			fmt.Printf("%s | 役种：", tt.HumanReadableString())
			if len(huway.YiZhongs) == 0 {
				fmt.Print("无役(0)，")
			}
			for _, yizhong := range huway.YiZhongs {
				fmt.Print(yizhong.GetName())
				if yizhong.(ruleset_jp.JapaneseMahjongYiZhong).IsYakuman() {
					fmt.Print("(役满)，")
				} else {
					fmt.Printf("(%d番)，", yizhong.GetFan())
				}
			}
			fmt.Print("| ")
			for _, fulu := range hand.Fulu {
				fmt.Printf("%s，", fulu.HumanReadableString())
			}
			fmt.Printf(" | %s\n", huway.HumanReadableString())
		}
	}
}
