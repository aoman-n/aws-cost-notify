package client

import (
	"aws-billing-notify/model"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

type AwsClient struct {
	cloudwatchSrv *cloudwatch.CloudWatch
	costexplorerSrv *costexplorer.CostExplorer
}

var _ AwsCostFetcher = (*AwsClient)(nil)

func NewAwsClient() *AwsClient {
	sess := session.Must(session.NewSession())
	cloudwatchSrv := cloudwatch.New(
		sess,
		aws.NewConfig().WithRegion("us-east-1"),
	)
	costexplorerSrv := costexplorer.New(sess)

	return &AwsClient{
		cloudwatchSrv: cloudwatchSrv,
		costexplorerSrv: costexplorerSrv,
	}
}

const (
	dateFormatForSDK = "2006-01-02"
)

func (a *AwsClient) FetchCost(startDate, endDate time.Time) (*model.AwsCost, error) {
	metric := costexplorer.MetricNetUnblendedCost
	metrics := []*string{&metric}
	beforDateStr := startDate.Format(dateFormatForSDK)
	targetDateStr := endDate.Format(dateFormatForSDK)
	timePeriod := costexplorer.DateInterval{
        Start: &beforDateStr,
        End:   &targetDateStr,
    }

	inputParams := &costexplorer.GetCostAndUsageInput{
        Granularity: aws.String(costexplorer.GranularityDaily),
        Metrics:     metrics,
        TimePeriod:  &timePeriod,
    }
	ret, err := a.costexplorerSrv.GetCostAndUsage(inputParams)
	if err != nil {
		return nil, err
	}

	amount := ret.ResultsByTime[0].Total["NetUnblendedCost"].Amount

	return &model.AwsCost{
		StartDate: startDate,
		EndDate: endDate,
		Amount: *amount,
	}, nil
}


// // fetchMonthCost 引数で渡された日のまでの月間コストを取得。1日の場合は前月のコストを取得
// func (a *AwsClient) fetchMonthCost(date time.Time) (string, error) {
// 	return "", nil
// }

// // retrun format: "2006-01-02"
// func formatDate(date time.Time) *string {
// 	formattedDate := date.Format(time.RFC3339)
// 	formattedDate = formattedDate[:strings.Index(formattedDate, "T")]
// 	return &formattedDate
// }

// // fetchDayCost 引数で渡された日の料金を取得
// func (a *AwsClient) fetchDayCost(targetDate time.Time) (string, error) {
// 	beforeDate := targetDate.AddDate(0, 0, -1)
// 	cost, err := a.fetchCost(beforeDate, targetDate)
// 	if err != nil {
// 		return "", err
// 	}

// 	return fmt.Sprintf("%s: $%s", targetDate.Format(dateFormatForOutput), cost), nil
// }

// fetchDayCost 引数で渡された月の料金を取得
// func (a *AwsClient) FetchCost(month time.Month) (*model.AwsCost, error) {
// 	period := "Monthly"
// 	// 現在時刻の取得
//     jst, _ := time.LoadLocation("Asia/Tokyo")
//     now := time.Now().UTC().In(jst)
//     dayBefore := now.AddDate(0, 0, -1)
//     first := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, jst)
//     if now.Day() == 1 { // 月初のときは先月分
// 		first = first.AddDate(0, -1, 0)
//     }
//     nowDate := now.Format("2006-01-02")
//     nowDateP := &nowDate
//     dateBefore := dayBefore.Format("2006-01-02")
//     dateBeforeP := &dateBefore
//     firstDate := first.Format("2006-01-02")
//     firstDateP := &firstDate

//     start := dateBeforeP
//     if period == "Monthly" {
//         start = firstDateP
//     }

//     granularity := costexplorer.GranularityDaily
//     if period == "Monthly" {
//         granularity = costexplorer.GranularityMonthly
//     }

//     // metric := "NetUnblendedCost" // 非ブレンド純コスト
//     metric := costexplorer.MetricNetUnblendedCost // 非ブレンド純コスト
//     metrics := []*string{&metric}

//     timePeriod := costexplorer.DateInterval{
//         Start: start,
//         End:   nowDateP,
//     }

//     // Inputの作成
//     input := &costexplorer.GetCostAndUsageInput{
//         Granularity: aws.String(granularity),
//         Metrics:     metrics,
//         TimePeriod:  &timePeriod,
//     }

//     // 処理実行
//     result, err := a.costexplorerSrv.GetCostAndUsage(input)
//     if err != nil {
//         log.Println(err.Error())
//     }

//     // 処理結果を出力
// 	log.Println(result)

// 	s := result.String()

// 	log.Println("s: ", s)

// 	// t := result.ResultsByTime[0].Total
// 	// amount := t["NetUnblendedCost"].Amount

// 	// total := result.ResultsByTime[0].Total
// 	// totalCostString := fmt.Sprintf(
// 	// 	"%s %s",
// 	// 	*total[costexplorer.MetricNetUnblendedCost].Unit,
// 	// 	*total[costexplorer.MetricNetUnblendedCost].Amount,
// 	// )

// 	// fmt.Println("totalCostString: ", totalCostString)

// 	return nil, nil
// }
