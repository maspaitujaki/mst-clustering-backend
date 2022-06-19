package entity

import "time"

type Log struct {
	Id              int       `json:"id"`
	NamaFile        string    `json:"nama_file"`
	HasilClustering string    `json:"hasil_clustering"`
	Tanggal         time.Time `json:"tanggal"`
}
