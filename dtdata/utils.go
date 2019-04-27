package dtdata

import (
	"crypto/md5"
	"fmt"
	"sort"

	"github.com/zhs007/dtdataserv/jarviscrawlercore"
	"github.com/zhs007/dtdataserv/proto"
	"github.com/zhs007/jarviscore"
)

func findBusinessInDTGameReport(game *dtdatapb.DTGameReport, businessid string) *dtdatapb.DTBusinessReport {
	for _, v := range game.BusinessReport {
		if v.BusinessID == businessid {
			return v
		}
	}

	return nil
}

func hasGameInDTBusinessReport(business *dtdatapb.DTBusinessReport, gameCode string) bool {
	for _, v := range business.GameReport {
		if v.GameCode == gameCode {
			return true
		}
	}

	return false
}

func findDTGameReport(lstGame []*dtdatapb.DTGameReport, gameCode string) *dtdatapb.DTGameReport {
	for _, v := range lstGame {
		if v.GameCode == gameCode {
			return v
		}
	}

	return nil
}

func findDTBusinessReport(lstBusiness []*dtdatapb.DTBusinessReport, businessid string) *dtdatapb.DTBusinessReport {
	for _, v := range lstBusiness {
		if v.BusinessID == businessid {
			return v
		}
	}

	return nil
}

func countDTReportWithBusinessGameReport(reply *jarviscrawlercore.ReplyDTData, mainCurrency string, scale int) *dtdatapb.DTReport {

	dtreport := &dtdatapb.DTReport{
		MainCurrency: mainCurrency,
	}

	var lstGame []*dtdatapb.DTGameReport
	var lstBusiness []*dtdatapb.DTBusinessReport

	for _, v := range reply.GameReports {
		if v.Currency == mainCurrency {
			dtreport.TotalBet += v.TotalBet
			dtreport.TotalWin += v.TotalWin
			dtreport.SpinNums += int64(v.GameNums)

			cg := findDTGameReport(lstGame, v.Gamecode)
			if cg == nil {
				cg = &dtdatapb.DTGameReport{
					GameCode: v.Gamecode,
				}

				lstGame = append(lstGame, cg)
			}

			cg.TotalBet += v.TotalBet
			cg.TotalWin += v.TotalWin
			cg.SpinNums += int64(v.GameNums)

			cgb := findBusinessInDTGameReport(cg, v.Businessid)
			if cgb == nil {
				cg.BusinessReport = append(cg.BusinessReport, &dtdatapb.DTBusinessReport{
					BusinessID: v.Businessid,
					TotalBet:   v.TotalBet,
					TotalWin:   v.TotalWin,
					SpinNums:   int64(v.GameNums),
				})

				cg.BusinessNums++
			} else {
				cgb.TotalBet += v.TotalBet
				cgb.TotalWin += v.TotalWin
				cgb.SpinNums += int64(v.GameNums)
			}

			cb := findDTBusinessReport(lstBusiness, v.Businessid)
			if cb == nil {
				cb = &dtdatapb.DTBusinessReport{
					BusinessID: v.Businessid,
				}

				lstBusiness = append(lstBusiness, cb)
			}

			cb.TotalBet += v.TotalBet
			cb.TotalWin += v.TotalWin
			cb.SpinNums += int64(v.GameNums)

			if !hasGameInDTBusinessReport(cb, v.Gamecode) {
				cb.GameReport = append(cb.GameReport, &dtdatapb.DTGameReport{
					GameCode: v.Gamecode,
					TotalBet: v.TotalBet,
					TotalWin: v.TotalWin,
					SpinNums: int64(v.GameNums),
				})
				// cb.Gamecode = append(cb.Gamecode, v.Gamecode)

				cb.GameNums++
			}
		}
	}

	dtreport.GameNums = int32(len(lstGame))
	dtreport.BusinessNums = int32(len(lstBusiness))

	sort.Slice(lstGame, func(i, j int) bool {
		return lstGame[i].TotalBet > lstGame[j].TotalBet
	})

	sort.Slice(lstBusiness, func(i, j int) bool {
		return lstBusiness[i].TotalBet > lstBusiness[j].TotalBet
	})

	for i := 0; i < len(lstGame); i++ {
		// for _, v := range lstGame[i].Businessid {
		// 	ccb := findDTBusinessReport(lstBusiness, v)
		// 	if ccb != nil {
		// 		lstGame[i].BusinessReport = append(lstGame[i].BusinessReport, &plugindtdatapb.DTBusinessReport{
		// 			BusinessID: ccb.G
		// 		})
		// 	}
		// }

		// lstGame[i].Businessid = nil

		dtreport.TopGames = append(dtreport.TopGames, lstGame[i])
	}

	for i := 0; i < len(lstBusiness); i++ {
		// for _, v := range lstBusiness[i].Gamecode {
		// 	ccg := findDTGameReport(lstGame, v)
		// 	if ccg != nil {
		// 		lstBusiness[i].GameReport = append(lstBusiness[i].GameReport, ccg)
		// 	}
		// }

		// lstBusiness[i].Gamecode = nil

		dtreport.TopBusiness = append(dtreport.TopBusiness, lstBusiness[i])
	}

	if scale > 0 && dtreport.TotalBet > 0 {
		maxbet := dtreport.TotalBet

		dtreport.TotalBet = float32(scale)
		dtreport.TotalWin = float32(scale) * dtreport.TotalWin / maxbet

		for _, v := range dtreport.TopGames {
			v.TotalBet = float32(scale) * v.TotalBet / maxbet
			v.TotalWin = float32(scale) * v.TotalWin / maxbet

			for _, vb := range v.BusinessReport {
				vb.TotalBet = float32(scale) * vb.TotalBet / maxbet
				vb.TotalWin = float32(scale) * vb.TotalWin / maxbet
			}
		}

		for _, v := range dtreport.TopBusiness {
			v.TotalBet = float32(scale) * v.TotalBet / maxbet
			v.TotalWin = float32(scale) * v.TotalWin / maxbet

			for _, vg := range v.GameReport {
				vg.TotalBet = float32(scale) * vg.TotalBet / maxbet
				vg.TotalWin = float32(scale) * vg.TotalWin / maxbet
			}
		}
	}

	return dtreport
}

// buildHashBuffer - build hash buffer
func buildHashBuffer(buf []byte) string {
	h := md5.New()
	return fmt.Sprintf("%x", h.Sum(buf))
}

// makeBusinessDayDataKey - make business day data key
func makeBusinessDayDataKey(env string, daytime string) string {
	return jarviscore.AppendString(env, ":bdd:", daytime)
}

// makeCacheKey - make cache key
func makeCacheKey(hashstr string) string {
	return jarviscore.AppendString("cache:", hashstr)
}

// newErrorDTReport - new DTReport with error
func newErrorDTReport(err error) *dtdatapb.DTReport {
	return &dtdatapb.DTReport{
		Err: err.Error(),
	}
}
