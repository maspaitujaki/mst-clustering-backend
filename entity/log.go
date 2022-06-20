package entity

import "time"

type Log struct {
	Id           int       `json:"id"`
	NamaFileAsli string    `json:"nama_file_asli"`
	Tanggal      time.Time `json:"tanggal"`
	N_Cluster    int       `json:"n_cluster"`
}
