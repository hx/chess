package chess

import "regexp"

func (r MoveRequest) UCI() string {
	str := r.From.String() + r.To.String()
	if r.PromoteTo != 0 {
		str += string(r.PromoteTo.Letter())
	}
	return str
}

var uciPattern = regexp.MustCompile(`(?i)^([a-h][1-8])([a-h][1-8])([qrnb]?)$`)

func ParseUCI(uci string) (MoveRequest, error) {
	match := uciPattern.FindStringSubmatch(uci)
	if match == nil {
		return MoveRequest{}, ErrInvalidMoveRequest
	}
	req := MoveRequest{
		From: MustParseCoordinate(match[1]),
		To:   MustParseCoordinate(match[2]),
	}
	if len(match[3]) == 1 {
		req.PromoteTo = MustParsePieceTypeLetter(rune(match[3][0]))
	}
	return req, nil
}

func MustParseUCI(uci string) MoveRequest {
	r, err := ParseUCI(uci)
	if err != nil {
		panic(err)
	}
	return r
}
