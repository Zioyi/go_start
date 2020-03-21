package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

type Vec2 struct {
	X, Y float32
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOveMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOveMag, v.Y * oneOveMag}
	}

	return Vec2{0, 0}
}

func (v Vec2) Draw(pic *image.Gray, x_offset int, y_offset int) {
	// 坐标从左上角开始
	pic.SetGray(int(v.X*10)+x_offset, y_offset-int(v.Y*10), color.Gray{0})

}

//
type Player struct {
	currPos   Vec2
	targetPos Vec2
	speed     float32
}

func (p *Player) MoveTo(v Vec2) {

	p.targetPos = v
}

func (p *Player) Pos() Vec2 {

	return p.currPos
}

func (p *Player) IsArrive() bool {

	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

func (p *Player) Update() {

	if !p.IsArrive() {

		// 计算当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()

		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))

		p.currPos = newPos
	}
}

func NewPlayer(speed float32) *Player {

	return &Player{
		speed: speed,
	}
}

func main() {

	const size = 100

	pic := image.NewGray(image.Rect(0, 0, size, size))

	var gray_v uint8 = 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if x == 0 || (y == 100) {
				gray_v = 0
			} else {
				gray_v = 255
			}
			pic.SetGray(x, y, color.Gray{gray_v})
		}
	}

	p := NewPlayer(0.5)

	p.MoveTo(Vec2{3, 1})
	// pic.SetGray(32, 12, color.Gray{0})

	for !p.IsArrive() {

		p.Update()

		currPos := p.Pos()
		fmt.Println(currPos)
		currPos.Draw(pic, size, 0, 100)

		// pic.SetGray(int(currPos.X*10)+2, int(currPos.Y*10)+2, color.Gray{0})

	}

	file, err := os.Create("/Users/zioyi/go/src/go_start/chapter05/player_trace.png")

	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, pic)
	file.Close()
}
