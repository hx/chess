package chess

type Outcome string

const (
	OutcomeNone                 Outcome = "Ongoing"
	OutcomeCheckmate            Outcome = "Checkmate"
	OutcomeStalemate            Outcome = "Stalemate"
	OutcomeInsufficientMaterial Outcome = "Insufficient material"
)

func (g *Game) IsEnded() bool {
	return g.IsMate() || g.isInsufficientMaterial()
}

func (g *Game) IsCheck() bool {
	return g.IsStarted() && g.lastMove().IsCheck
}

func (g *Game) IsMate() bool {
	return len(g.legalMoves) == 0
}

func (g *Game) IsCheckmate() bool {
	return g.IsCheck() && g.IsMate()
}

func (g *Game) IsStalemate() bool {
	return g.IsMate() && !g.IsCheck()
}

func (g *Game) IsInsufficientMaterial() bool {
	return !g.IsMate() && g.isInsufficientMaterial()
}

func (g *Game) Outcome() Outcome {
	if g.IsMate() {
		if g.IsCheck() {
			return OutcomeCheckmate
		}
		return OutcomeStalemate
	}
	if g.isInsufficientMaterial() {
		return OutcomeInsufficientMaterial
	}
	return OutcomeNone
}
