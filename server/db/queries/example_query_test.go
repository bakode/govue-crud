package queries

import "testing"

func TestAllData(t *testing.T) {
	var total = 2 // this is total data in database
	t.Log("Data size : %D", len(All()))

	if len(All()) != total {
		t.Errorf("jumlah total data tidak sama")
	}
}

func TestFindData(t *testing.T) {
	var id = 2
	t.Log("data title: %S", Find(id).Title)

	if Find(id).Id != 2 {
		t.Errorf("sepertinya id salah")
	}
}

func TestStoreData(t *testing.T) {
	t.Log("Test Store Data")

	if Store("Baru", "Dicoba baru") != true {
		t.Errorf("gagal tambah data")
	}
}

func TestUpdateData(t *testing.T) {
	t.Log("Test Update data")
	var id = 7
	if Update(id , "Update", "Dicoba update") != true {
		t.Errorf("gagal update data")
	}
}

func TestDestroyData(t *testing.T) {
	t.Log("Test Delete data")
	var id = 2
	if Destroy(id) != true {
		t.Errorf("gagal delete data")
	}
}