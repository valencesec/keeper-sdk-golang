package enterprise

import (
	"github.com/keeper-security/keeper-sdk-golang/auth"
	"github.com/keeper-security/keeper-sdk-golang/proto_auth"
)

type SecurityReportEntry struct {
	EnterpriseUserId       int64
	LastLogin              int64
	TwoFactor              string
	NumberOfReusedPassword int32
	UserId                 int32
}

func GetSecurityReport(keeperAuth auth.IKeeperAuth) ([]SecurityReportEntry, error) {
	var entries []SecurityReportEntry
	var fromPage int64

	for {
		rq := &proto_auth.SecurityReportRequest{FromPage: fromPage}
		rs := &proto_auth.SecurityReportResponse{}
		if err := keeperAuth.ExecuteAuthRest("enterprise/get_security_report_data", rq, rs); err != nil {
			return nil, err
		}

		for _, r := range rs.SecurityReport {
			entries = append(entries, SecurityReportEntry{
				EnterpriseUserId:       r.GetEnterpriseUserId(),
				LastLogin:              r.GetLastLogin(),
				TwoFactor:              r.GetTwoFactor(),
				NumberOfReusedPassword: r.GetNumberOfReusedPassword(),
				UserId:                 r.GetUserId(),
			})
		}

		if rs.Complete {
			break
		}
		fromPage = rs.ToPage
	}

	return entries, nil
}
