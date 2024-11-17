package frames

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Scene struct {
	Screen        [][]rune
	InitialScreen [][]rune
	changed       bool
}

const emptySpace = " "
const penisStartLine = 5

func NewScene(frame string) *Scene {
	return &Scene{
		Screen:        splitIntoLines(frame),
		InitialScreen: splitIntoLines(frame),
	}
}
func splitIntoLines(frame string) [][]rune {
	lines := make([][]rune, 0)
	s := strings.Split(frame, "\n")

	for _, line := range s {
		lines = append(lines, []rune(line))
	}

	return lines
}

func (s *Scene) addOffsetToLine(line, offset int) {
	s.Screen[line] = append(s.Screen[line], []rune(strings.Repeat(emptySpace, offset))...)
}

func (s *Scene) toString() string {
	str := ""
	for _, line := range s.Screen {
		str += fmt.Sprintf("%s\n", string(line))
	}
	return str
}
func calcOffset(prev, curr, step int) int {
	if curr > prev {
		return curr - prev + step
	}

	return prev - curr + step
}
func (s *Scene) resetLine(line int) {
	s.Screen[line] = splitIntoLines(MainCharacter)[line]
	s.changed = true
}

func configurePenisBullet(bullet string, bulletCount int, penisBulletOffset int) string {
	str := ""
	for i := 0; i < bulletCount; i++ {
		str += strings.Repeat(emptySpace, penisBulletOffset)
		str += bullet
	}

	return str
}
func (s *Scene) addShaking(offset int) {
	for i := 0; i < len(s.Screen)-2; i++ {
		newLine := []rune(strings.Repeat(emptySpace, offset))
		newLine = append(newLine, s.Screen[i]...)
		s.Screen[i] = newLine
	}
	s.changed = true
	time.Sleep(time.Millisecond * 100)
	for i := 0; i < len(s.Screen)-2; i++ {
		s.resetLine(i)
	}
}
func setMin(v int) int {
	if v < 5 {
		return 5 + v
	}

	return v
}
func (s *Scene) animateFlySusSubstantion(susSubstantion string) {
	step := setMin(int(rand.Int31n(10)))
	start := len([]rune(s.Screen[penisStartLine]))
	s.addShaking(1)
	for i := penisStartLine; i < len(s.Screen)-2; i++ {
		if i != penisStartLine {
			s.addOffsetToLine(i, calcOffset(start, len(s.Screen[i]), step))
		}
		penisBullet := []rune(configurePenisBullet(susSubstantion, 1, 1))
		for j := 0; j < len(penisBullet); j++ {
			s.Screen[i] = append(s.Screen[i], penisBullet[j])
			s.changed = true
			time.Sleep(time.Millisecond * 7)
		}
		start += step
		time.Sleep(time.Millisecond * 20)
		s.resetLine(i)
	}
}

func (s *Scene) useEffect() {
	fmt.Print(s.toString())
	for {
		if s.changed {
			fmt.Print("\033[H\033[J")
			fmt.Print(s.toString())
			s.changed = false
		}
	}
}

func (s *Scene) Render() {
	go s.useEffect()
	for {
		s.animateFlySusSubstantion(SusSubstantion)
		s.Screen = splitIntoLines(MainCharacter)
		s.changed = true
	}
}
