package main

import (
	"fmt"
	"time"

	"github.com/0x5ab/gomj/gameplay"
	"github.com/0x5ab/gomj/ruleset"
	"github.com/0x5ab/gomj/ruleset/ruleset_jp"
	"github.com/0x5ab/gomj/tiles"
	"github.com/0x5ab/gomj/utils"
	"github.com/0x5ab/gomj/wind"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	// t, err := tiles.ParseTiles("1z1z1z2z2z2z5z")
	// t, err := tiles.ParseTiles("3m3m3m4m5m5m6m6m7m7m8m8m8m")
	t, err := tiles.ParseTiles("1112345678999m")
	// t, err := tiles.ParseTiles("2233445566778p")
	// t, err := tiles.ParseTiles("34888s66z")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)
	game := &gameplay.Game{Wind: wind.East, DrawsRemaining: 123}
	player := &gameplay.Player{Id: 3, Wind: wind.East, Game: game}
	// hand := &gameplay.Hand{
	// 	Game:   game,
	// 	Player: player,
	// 	Tiles:  t,
	// 	Fulu:   []gameplay.Fulu{{StartTile: tiles.Suo2, Type: gameplay.Chi}, {StartTile: tiles.Suo6, Type: gameplay.Peng}},
	// }
	hand := gameplay.NewHand(game, player)
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
		tile := &gameplay.GameTile{Tile: tt, Player: player}
		// tile := &gameplay.GameTile{Tile: tiles.Suo2, Player: player}
		huway := ruleset.CanHu(ruleset_jp.JapaneseMahjongRuleset, hand, tile)
		if !huway.IsValid() {
			continue
		}
		fmt.Printf("%s | 役种：", tt.HumanReadableString())
		if len(huway.YiZhongs) == 0 {
			fmt.Print("无役(0)，")
		}
		for _, yizhong := range huway.YiZhongs {
			fmt.Print(yizhong.GetName())
			if yizhong.(ruleset_jp.JapaneseMahjongYaku).YakumanBaisu() > 0 {
				if yizhong.(ruleset_jp.JapaneseMahjongYaku).YakumanBaisu() == 1 {
					fmt.Print("(役满)，")
				} else {
					fmt.Printf("(%s倍役满)，", utils.DigitToChinese(yizhong.(ruleset_jp.JapaneseMahjongYaku).YakumanBaisu()))
				}
			} else {
				fmt.Printf("(%d番)，", yizhong.GetFan(huway))
			}
		}
		fmt.Printf("| %d符", huway.Point)
		for _, fulu := range hand.Fulu {
			fmt.Printf(" | %s", fulu.HumanReadableString())
		}
		fmt.Printf(" | %s\n", huway.HumanReadableString())
	}
}
