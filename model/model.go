package model

type Then struct {
	Text string
}

type When struct {
	Text string
	Then []Then
}

type Given struct {
	Text string
	When []When
}
