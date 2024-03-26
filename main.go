package main

import (
	// "context"
	"embed"
	// "encoding/json"
	// "net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// WailsApp struct để giao tiếp với VueJS
// type WailsApp struct {
// 	deThi DeThi `json:"deThi"` // Lưu trữ dữ liệu đề thi
// }

// DeThi struct đại diện cho dữ liệu JSON đã phân tích
// type DeThi struct {
// 	Metadata struct {
// 		TenDeThi       string `json:"ten_de_thi"`
// 		ThoiGianLamBai int    `json:"thoi_gian_lam_bai"`
// 		SoLuongCauHoi  int    `json:"so_luong_cau_hoi"`
// 	} `json:"metadata"`
// 	DanhSachCauHoi []struct {
// 		NoiDungCauHoi string   `json:"noi_dung_cau_hoi"`
// 		LuaChon       []string `json:"lua_chon"`
// 		DapAnDung     string   `json:"dap_an_dung"`
// 	} `json:"danh_sach_cau_hoi"`
// }

// func startup(ctx context.Context) {
// }

// func handleAPI(w http.ResponseWriter, r *http.Request) {
// 	// Thiết lập tiêu đề Content-Type để chỉ định dữ liệu là JSON
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	// Tạo dữ liệu đề thi trực tiếp trong mã
// 	deThi := DeThi{
// 		Metadata: struct {
// 			TenDeThi       string `json:"ten_de_thi"`
// 			ThoiGianLamBai int    `json:"thoi_gian_lam_bai"`
// 			SoLuongCauHoi  int    `json:"so_luong_cau_hoi"`
// 		}{
// 			TenDeThi:       "Đề thi thử môn Toán",
// 			ThoiGianLamBai: 60,
// 			SoLuongCauHoi:  10,
// 		},
// 		DanhSachCauHoi: []struct {
// 			NoiDungCauHoi string   `json:"noi_dung_cau_hoi"`
// 			LuaChon       []string `json:"lua_chon"`
// 			DapAnDung     string   `json:"dap_an_dung"`
// 		}{
// 			{
// 				NoiDungCauHoi: "Câu 1: Căn bậc hai của 4 là gì?",
// 				LuaChon:       []string{"1", "2", "3", "4"},
// 				DapAnDung:     "2",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 2: Căn bậc hai của 9 là gì?",
// 				LuaChon:       []string{"1", "2", "3", "4"},
// 				DapAnDung:     "3",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 3: Khái niệm hệ điều hành là gì?",
// 				LuaChon: []string{"Cung cấp và xử lý các phần cứng và phần mềm.",
// 					"Nghiên cứu phương pháp, kỹ thuật xử lý thông tin bằng máy tính điện tử.",
// 					"Nghiên cứu về công nghệ phần cứng và phần mềm",
// 					"Là một phần mềm chạy trên máy tính, dùng để điều hành, quản lý các thiết bị phần cứng và các tài nguyên phần mềm trên máy tính"},
// 				DapAnDung: "Là một phần mềm chạy trên máy tính, dùng để điều hành, quản lý các thiết bị phần cứng và các tài nguyên phần mềm trên máy tính",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 4: Phần mềm nào gọi là hệ điều hành?",
// 				LuaChon:       []string{"Windows 10", "Unikey", "Microsoft Word", "Microsoft Excel"},
// 				DapAnDung:     "Windows 10",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 5: Phần mềm giao tiếp thiết bị gọi là:",
// 				LuaChon:       []string{"Driver", "Hệ điều hành", "Thiết bị", "Bus"},
// 				DapAnDung:     "Driver",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 6: Hệ điều hành chạy khi nào?",
// 				LuaChon:       []string{"Sau khi bật máy tính", "Khi người sử dụng click chuột", "Khi tắt máy", "Khi người sử dụng Enter"},
// 				DapAnDung:     "Sau khi bật máy tính",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 7: Hệ điều hành cài đặt trên:",
// 				LuaChon:       []string{"Đĩa cứng", "RAM", "ROM", "ROM"},
// 				DapAnDung:     "Đĩa cứng",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 8: Những hệ điều hành nào của Microsoft?",
// 				LuaChon:       []string{"Windows", "Unix", "Linix", "Ubuntu"},
// 				DapAnDung:     "Windows",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 9: Những hệ điều hành nào có mã nguồn mở?",
// 				LuaChon:       []string{"Windows", "MS-DOS", "Unix", "IOS"},
// 				DapAnDung:     "Unix",
// 			},
// 			{
// 				NoiDungCauHoi: "Câu 10: Hệ thống tập tin nào của hệ điều hành Linux?",
// 				LuaChon:       []string{"FAT16", "FAT32", "Ext", "NTFS"},
// 				DapAnDung:     "Ext",
// 			},
// 		},
// 	}
// 	w.Header().Set("Content-Type", "application/json")

// 	// Trả về dữ liệu đề thi dưới dạng JSON
// 	json.NewEncoder(w).Encode(deThi)
// }

func main() {
	// go func() {
	// 	http.HandleFunc("/api/de_thi", handleAPI)
	// 	http.ListenAndServe(":8080", nil)
	// }()

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
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
