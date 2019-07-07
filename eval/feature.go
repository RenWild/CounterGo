//go:generate stringer -type Feature

package eval

type Feature int

const (
	fPawnMaterial Feature = iota
	fKnightMaterial
	fBishopMaterial
	fRookMaterial
	fQueenMaterial
	fBishopPairMaterial
	fPawnWeak
	fPawnDoubled
	fPawnDuo
	fPawnProtected
	fPawnPassed
	fPawnPassedFree
	fPawnPassedOppKing
	fPawnPassedOwnKing
	fPawnPassedSquare
	fThreatPawn
	fThreatForPawn
	fThreatPiece
	fThreatPieceForQueen
	fKnightPst
	fQueenPst
	fKingCastlingPst
	fKingCenterPst
	fKnightMobility
	fBishopMobility
	fRookMobility
	fQueenMobility
	fRook7th
	fRookOpen
	fRookSemiopen
	fKingShelter
	fKingAttack
	fKingAttack2
	fKingQueenTropism
	fBishopRammedPawns
	fMinorProtected
	fKnightOutpost
	fPawnBlockedByOwnPiece
	fPawnRammed
	fSize
)

// Error: 0.055977
// Regularization: 0.002955
// Total: 0.058932
var autoGeneratedWeights = []int{9100, 11400, 33600, 31300, 36900, 34300, 48100, 59900, 118100, 118100, 4200, 5300, -1302, -930, -1600, -1600, 664, 0, 1100, 800, 234, 594, 0, 627, 0, 300, 12, -148, 0, 6400, 5900, 800, 0, 1800, 2400, 2100, 5100, 0, 682, 1023, -700, 0, -85, 255, -560, 735, 136, 36, 160, 152, 128, 212, 140, 310, 400, 900, 4400, -500, 1100, 100, 1521, -195, 4100, 500, 13900, 1700, -128, -2304, -700, -1300, 700, 2000, 2500, 300, -1800, 0, 0, -1100}
