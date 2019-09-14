package crunchyroll

// import (
// 	"hikari/players"
// 	"testing"
// )
//
// func TestParse(t *testing.T) {
// 	crTest := []struct {
// 		WTitle players.WindowTitle
// 		Expected bool
// 	}{
// 		{"Crunchyroll - Watch Popular Anime & Read Manga Online - Mozilla Firefox", false},
// 		{"Free Anime Streaming Online - Watch on Crunchyroll - Mozilla Firefox", false},
// 		{"Free Drama Streaming Online - Watch on Crunchyroll - Mozilla Firefox", false},
// 		{"Crunchyroll - Browse Simulcasted Anime - Mozilla Firefox", false},
// 		{"Crunchyroll - Browse Recently Upddated Anime - Mozilla Firefox", false},
// 		{"Crunchyroll - Browse Anime Alphabetically - Mozilla Firefox", false},
// 		{"Crunchyroll - Browse Anime By Genres - Mozilla Firefox", false},
// 		{"Crunchyroll - Simulcast Calendar - Mozilla Firefox", false},
// 		{"Crunchyroll - Spring 2019 Lineup - Mozilla Firefox", false},
// 		{"Crunchyroll - Crunchyroll-Ready Devices - Mozilla Firefox", false},
// 		{"Crunchyroll - Premium Membership Free Trial - Mozilla Firefox", false},
// 		{"Manga Lineup - Mozilla Firefox", false},
// 		{"Read Popular Manga Online - Crunchyroll - Mozilla Firefox", false},
// 		{"Crunchyroll - Crunchyroll Manga - Mozilla Firefox", false},
// 		{"Crunchyroll - Manga Title - Mozilla Firefox", false},
// 		{"Crunchyroll - News - Mozilla Firefox", false},
// 		{"Crunchyroll - Article Title - Mozilla Firefox", false},
// 		{"Crunchyroll - Username - Profile Info - Mozilla Firefox", false},
// 		{"Crunchyroll - Forum Main - Mozilla Firefox", false},
// 		{"Crunchyroll - Forum - Section - Mozilla Firefox", false},
// 		{"Crunchyroll - Forum - Thread title - Mozilla Firefox", false},
// 		{"Crunchyroll - Forum - General - Mozilla Firefox", false},
// 		{"Crunchyroll - Forum - Discussion - Mozilla Firefox", false},
// 		{"Crunchyroll - Forum - Crunchy Connect - Mozilla Firefox", false},
// 		{"Crunchyroll - Forum - Creative - Mozilla Firefox", false},
// 		{"Crunchyroll - Forum - Help - Mozilla Firefox", false},
// 		{"Crunchyroll - Library - Mozilla Firefox", false},
// 		{"Crunchyroll - Help - Contact Us - Mozilla Firefox", false},
// 		{"Crunchyroll - Help - Video Troubleshooting - Mozilla Firefox", false},
// 		{"Anime Merchandise: Figures, Shirts, DVDs & More | Crunchyroll Store - Mozilla Firefox", false},
// 		{"Crunchyroll - Mozilla Firefox", false},
// 		{"Crunchyroll - Sign up or Log In - Mozilla Firefox", false},
// 		{"Crunchyroll - Terms of Service - Mozilla Firefox", false},
// 		{"Crunchyroll - Privacy Policy - Mozilla Firefox", false},
// 		{"Crunchyroll - My Queue - Mozilla Firefox", false},
// 		// "Title - Watch on Crunchyroll - Mozilla Firefox",          -- can't really handle movies because of the regular title pages.
// 		{"Show Information - Crunchyroll - Mozilla Firefox", false},
// 		{"Comments - Crunchyroll - Mozilla Firefox", false},
// 		{"Reviews - Crunchyroll - Mozilla Firefox", false},
// 		{"Title PV 1 - Watch on Crunchyroll", false},
// 		{"The Rising of the Shield Hero Episode 12, The Raven Invader, - Watch on Crunchyroll - Mozilla Firefox", true},
// 		{"My Hero Academia Season 3 Episode 45, What a Twist!, - Watch on Crunchyroll - Mozilla Firefox", true},
// 		{"The Ancient Magus' Bride Episode 10, We live and learn, - Watch on Crunchyroll - Mozilla Firefox", true},
// 		{"The Ancient Magus' Bride: Those Awaiting a Star Episode 2, Part 2, - Watch on Crunchyroll - Mozilla Firefox", true}, // OVA
// 		{"Naturo Shippuden: The Past: The Hidden Leaf Village Episode 190, Naruto and the Old Soldier, - Watch on Crunchyroll - Mozilla Firefox", true},
// 		{"One Piece: Whole Cake Island (783-current) Episode 861, The Cake Sank?! Sanji and Bege's Getaway Battle!, - Watch on Crunchyroll - Mozilla Firefox", true},
// 		// "One Piece - Episode of East Blue Episode of East Blue - Watch on Crunchyroll - Mozilla Firefox",         -- can't really handle movies because of the regular title pages.
// 	}
// 	for _, v := range crTest {
// 		cr, _ := v.WTitle.GetMedia()
//
// 		if cr == nil && v.Expected == true {
// 			t.Error("Expected", v.Expected, "Got", false)
// 		} else if cr != nil && v.Expected == false {
// 			t.Error("Expected", v.Expected, "Got", true)
// 		}
// 	}
// }
