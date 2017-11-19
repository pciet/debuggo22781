// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"github.com/pciet/debuggo22781/wichessing"
)

const clients = 6

func main() {
	for i := 0; i < clients; i++ {
		go simulateClient()
	}
	select {} // block forever
}

func simulateClient() {
	for {
		board := wichessing.Board{}
		for i := 0; i < 64; i++ {
			board[i] = wichessing.Point{
				AbsPoint: wichessing.AbsPoint{wichessing.FileFromIndex(uint8(i)), wichessing.RankFromIndex(uint8(i))},
			}
		}
		setChessStart(&board)
		_, _, _ = board.Moves(wichessing.White, wichessing.AbsPoint{0, 8}, wichessing.AbsPoint{0, 8})
	}
}

func setChessStart(b *wichessing.Board) {
	b[0].Piece = &wichessing.Piece{
		Kind:        wichessing.Rook,
		Orientation: wichessing.White,
		Base:        wichessing.Rook,
	}
	b[1].Piece = &wichessing.Piece{
		Kind:        wichessing.Knight,
		Orientation: wichessing.White,
		Base:        wichessing.Knight,
	}
	b[2].Piece = &wichessing.Piece{
		Kind:        wichessing.Bishop,
		Orientation: wichessing.White,
		Base:        wichessing.Bishop,
	}
	b[3].Piece = &wichessing.Piece{
		Kind:        wichessing.Queen,
		Orientation: wichessing.White,
		Base:        wichessing.Queen,
	}
	b[4].Piece = &wichessing.Piece{
		Kind:        wichessing.King,
		Orientation: wichessing.White,
		Base:        wichessing.King,
	}
	b[5].Piece = &wichessing.Piece{
		Kind:        wichessing.Bishop,
		Orientation: wichessing.White,
		Base:        wichessing.Bishop,
	}
	b[6].Piece = &wichessing.Piece{
		Kind:        wichessing.Knight,
		Orientation: wichessing.White,
		Base:        wichessing.Knight,
	}
	b[7].Piece = &wichessing.Piece{
		Kind:        wichessing.Rook,
		Orientation: wichessing.White,
		Base:        wichessing.Rook,
	}
	for i := 0; i < 8; i++ {
		b[8+i].Piece = &wichessing.Piece{
			Kind:        wichessing.Pawn,
			Orientation: wichessing.White,
			Base:        wichessing.Pawn,
		}
	}
	b[63].Piece = &wichessing.Piece{
		Kind:        wichessing.Rook,
		Orientation: wichessing.Black,
		Base:        wichessing.Rook,
	}
	b[62].Piece = &wichessing.Piece{
		Kind:        wichessing.Knight,
		Orientation: wichessing.Black,
		Base:        wichessing.Knight,
	}
	b[61].Piece = &wichessing.Piece{
		Kind:        wichessing.Bishop,
		Orientation: wichessing.Black,
		Base:        wichessing.Bishop,
	}
	b[60].Piece = &wichessing.Piece{
		Kind:        wichessing.King,
		Orientation: wichessing.Black,
		Base:        wichessing.King,
	}
	b[59].Piece = &wichessing.Piece{
		Kind:        wichessing.Queen,
		Orientation: wichessing.Black,
		Base:        wichessing.Queen,
	}
	b[58].Piece = &wichessing.Piece{
		Kind:        wichessing.Bishop,
		Orientation: wichessing.Black,
		Base:        wichessing.Bishop,
	}
	b[57].Piece = &wichessing.Piece{
		Kind:        wichessing.Knight,
		Orientation: wichessing.Black,
		Base:        wichessing.Knight,
	}
	b[56].Piece = &wichessing.Piece{
		Kind:        wichessing.Rook,
		Orientation: wichessing.Black,
		Base:        wichessing.Rook,
	}
	for i := 0; i < 8; i++ {
		b[63-16+i].Piece = &wichessing.Piece{
			Kind:        wichessing.Pawn,
			Orientation: wichessing.Black,
			Base:        wichessing.Pawn,
		}
	}
	for i := 0; i < 64; i++ {
		if b[i].Piece == nil {
			continue
		}
		p := b[i].Piece.SetKindFlags()
		b[i].Piece = &p
	}
}
