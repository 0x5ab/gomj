package main

import (
	"fmt"

	"github.com/0x5ab/gomj/gameplay"
	"github.com/0x5ab/gomj/ruleset"
	"github.com/0x5ab/gomj/ruleset/ruleset_jp"
	"github.com/0x5ab/gomj/tiles"
	"github.com/0x5ab/gomj/wind"
)

func main() {
	// t, err := tiles.ParseTiles("1z1z1z2z2z2z5z")
	// t, err := tiles.ParseTiles("3m3m3m4m5m5m6m6m7m7m8m8m8m")
	t, err := tiles.ParseTiles("1m1m1m1m2m3m5m6m7m8m9m9m9m")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)
	// hand := gameplay.Hand{Tiles: t, Fulu: []gameplay.Fulu{{StartTile: tiles.Bai, Type: gameplay.Peng}, {StartTile: tiles.Fa, Type: gameplay.Peng}}}
	hand := gameplay.NewHand(&gameplay.Game{Wind: wind.East}, &gameplay.Player{Wind: wind.East})
	hand.Tiles = t
	cnt := make(map[int]int)
	for _, tile := range hand.Tiles {
		cnt[tile.Id]++
	}
	fmt.Printf("%s\n", hand.String())
	// x := ruleset.CanHu(&ruleset_jp.YiZhongRuleset, hand, &gameplay.GameTile{Tile: tiles.Wan9})
	// fmt.Printf("%+v\n", x)
	fmt.Println("能胡: ")
	for _, tt := range tiles.AllTiles {
		if cnt[tt.Id] >= 4 {
			continue
		}
		huway := ruleset.CanHu(ruleset_jp.JapaneseMahjongRuleset, hand, &gameplay.GameTile{Tile: tt})
		if !huway.IsValid() {
			continue
		}
		fmt.Printf("%s | 役种：", tt.HumanReadableString())
		if len(huway.YiZhongs) == 0 {
			fmt.Print("无役(0)，")
		}
		for _, yizhong := range huway.YiZhongs {
			fmt.Print(yizhong.GetName())
			if yizhong.(ruleset_jp.JapaneseMahjongYiZhong).IsYakuman() {
				fmt.Print("(役满)，")
			} else {
				fmt.Printf("(%d番)，", yizhong.GetFan(hand))
			}
		}
		fmt.Printf("| %d符", huway.Point)
		for _, fulu := range hand.Fulu {
			fmt.Printf("%s，", fulu.HumanReadableString())
		}
		fmt.Printf(" | %s\n", huway.HumanReadableString())
	}
}
