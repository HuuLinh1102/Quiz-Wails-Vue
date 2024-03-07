package main

import (
	"context"
	"embed"
	"encoding/json"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// WailsApp struct để giao tiếp với VueJS
type WailsApp struct {
	deThi DeThi `json:"deThi"` // Lưu trữ dữ liệu đề thi
}

// DeThi struct đại diện cho dữ liệu JSON đã phân tích
type DeThi struct {
	Metadata struct {
		TenDeThi       string `json:"ten_de_thi"`
		ThoiGianLamBai int    `json:"thoi_gian_lam_bai"`
		SoLuongCauHoi  int    `json:"so_luong_cau_hoi"`
	} `json:"metadata"`
	DanhSachCauHoi []struct {
		NoiDungCauHoi string   `json:"noi_dung_cau_hoi"`
		LuaChon       []string `json:"lua_chon"`
		DapAnDung     string   `json:"dap_an_dung"`
	} `json:"danh_sach_cau_hoi"`
}

func startup(ctx context.Context) {
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	// Thiết lập tiêu đề Content-Type để chỉ định dữ liệu là JSON
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Tạo dữ liệu đề thi trực tiếp trong mã
	deThi := DeThi{
		Metadata: struct {
			TenDeThi       string `json:"ten_de_thi"`
			ThoiGianLamBai int    `json:"thoi_gian_lam_bai"`
			SoLuongCauHoi  int    `json:"so_luong_cau_hoi"`
		}{
			TenDeThi:       "Đề thi thử môn Toán",
			ThoiGianLamBai: 60,
			SoLuongCauHoi:  10,
		},
		DanhSachCauHoi: []struct {
			NoiDungCauHoi string   `json:"noi_dung_cau_hoi"`
			LuaChon       []string `json:"lua_chon"`
			DapAnDung     string   `json:"dap_an_dung"`
		}{
			{
				NoiDungCauHoi: "Căn bậc hai của 4 là gì?",
				LuaChon:       []string{"1", "2", "3", "4"},
				DapAnDung:     "2",
			},
			{
				NoiDungCauHoi: "Căn bậc hai của 9 là gì?",
				LuaChon:       []string{"1", "2", "3", "4"},
				DapAnDung:     "3",
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")

	// Trả về dữ liệu đề thi dưới dạng JSON
	json.NewEncoder(w).Encode(deThi)
}

func main() {
	go func() {
		http.HandleFunc("/api/de_thi", handleAPI)
		http.ListenAndServe(":8080", nil)
	}()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "quiz",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
