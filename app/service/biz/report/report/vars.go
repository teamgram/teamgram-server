// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package report

import (
	"github.com/teamgram/proto/mtproto"
)

/*
# ReportReason
Report reason

```
inputReportReasonSpam#58dbcab8 = ReportReason;
inputReportReasonViolence#1e22c78d = ReportReason;
inputReportReasonPornography#2e59d922 = ReportReason;
inputReportReasonChildAbuse#adf44ee3 = ReportReason;
inputReportReasonOther#e1746d0a text:string = ReportReason;
inputReportReasonCopyright#9b89f93a = ReportReason;
inputReportReasonGeoIrrelevant#dbd4feed = ReportReason;
```

## Constructors
| Constructor |	Description |
| ----------- | ----------- |
| inputReportReasonSpam | Report for spam |
| inputReportReasonViolence | Report for violence |
| inputReportReasonPornography | Report for pornography |
| inputReportReasonChildAbuse | Report for child abuse |
| inputReportReasonOther | Other |
| inputReportReasonCopyright | Report for copyrighted content |
| inputReportReasonGeoIrrelevant | Report an irrelevant geogroup |
*/

type ReportReasonType int8

const (
	REASON_NONE           ReportReasonType = 0 // Unknown
	REASON_SPAM           ReportReasonType = 1 // Report for spam
	REASON_VIOLENCE       ReportReasonType = 2 // Report for violence
	REASON_PORNOGRAPHY    ReportReasonType = 3 // Report for pornography
	REASON_OTHER          ReportReasonType = 4 // Other
	REASON_COPYRIGHT      ReportReasonType = 5 // Report for copyrighted content
	REASON_CHILD_ABUSED   ReportReasonType = 6 // Report for child abuse
	REASON_GEO_IRRELEVANT ReportReasonType = 7 // Report an irrelevant geogroup
	REASON_FAKE           ReportReasonType = 8 // Report fake
)

/*
## Dealing with spam and ToS violations
| Name | Description |
| ---- | ----------- |
| account.reportPeer | Report a peer for violation of telegram's Terms of Service |
| channels.reportSpam | Reports some messages from a user in a supergroup as spam; requires administrator rights in the supergroup |
| messages.report | Report a message in a chat for violation of telegram's Terms of Service |
| messages.reportSpam | Report a new incoming chat for spam, if the peer settings of the chat allow us to do that |
| messages.reportEncryptedSpam | Report a secret chat for spam |
*/
const (
	ACCOUNTS_reportPeer          = 0
	MESSAGES_reportSpam          = 1
	MESSAGES_report              = 2
	MESSAGES_reportEncryptedSpam = 3
	CHANNELS_reportSpam          = 4
	ACCOUNTS_reportProfilePhoto  = 5
)

func (i ReportReasonType) String() (s string) {
	switch i {
	case REASON_SPAM:
		s = "inputReportReasonSpam#58dbcab8 = ReportReason"
	case REASON_VIOLENCE:
		s = "inputReportReasonPornography#2e59d922 = ReportReason"
	case REASON_PORNOGRAPHY:
		s = "inputReportReasonOther#e1746d0a text:string = ReportReason"
	case REASON_OTHER:
		s = "inputReportReasonOther#e1746d0a text:string = ReportReason"
	case REASON_COPYRIGHT:
		s = "inputReportReasonCopyright#9b89f93a = ReportReason;"
	case REASON_CHILD_ABUSED:
		s = "inputReportReasonChildAbuse#adf44ee3 = ReportReason;"
	case REASON_GEO_IRRELEVANT:
		s = "inputReportReasonGeoIrrelevant#dbd4feed = ReportReason;"
	case REASON_FAKE:
		s = "inputReportReasonFake#f5ddd6e7 = ReportReason;"
	default:
		s = "unknown"
	}
	return
}

func FromReportReason(reason *mtproto.ReportReason) (i ReportReasonType) {
	switch reason.PredicateName {
	case mtproto.Predicate_inputReportReasonSpam:
		i = REASON_SPAM
	case mtproto.Predicate_inputReportReasonViolence:
		i = REASON_VIOLENCE
	case mtproto.Predicate_inputReportReasonPornography:
		i = REASON_PORNOGRAPHY
	case mtproto.Predicate_inputReportReasonChildAbuse:
		i = REASON_CHILD_ABUSED
	case mtproto.Predicate_inputReportReasonOther:
		i = REASON_OTHER
	case mtproto.Predicate_inputReportReasonCopyright:
		i = REASON_COPYRIGHT
	case mtproto.Predicate_inputReportReasonGeoIrrelevant:
		i = REASON_GEO_IRRELEVANT
	case mtproto.Predicate_inputReportReasonFake:
		i = REASON_FAKE
	default:
		i = REASON_NONE
	}
	return
}

func (i ReportReasonType) ToReportReason(text string) (reason *mtproto.ReportReason) {
	switch i {
	case REASON_SPAM:
		reason = mtproto.MakeTLInputReportReasonSpam(nil).To_ReportReason()
	case REASON_VIOLENCE:
		reason = mtproto.MakeTLInputReportReasonViolence(nil).To_ReportReason()
	case REASON_PORNOGRAPHY:
		reason = mtproto.MakeTLInputReportReasonPornography(nil).To_ReportReason()
	case REASON_CHILD_ABUSED:
		reason = mtproto.MakeTLInputReportReasonChildAbuse(nil).To_ReportReason()
	case REASON_OTHER:
		reason = mtproto.MakeTLInputReportReasonOther(nil).To_ReportReason()
	case REASON_COPYRIGHT:
		reason = mtproto.MakeTLInputReportReasonCopyright(nil).To_ReportReason()
	case REASON_GEO_IRRELEVANT:
		reason = mtproto.MakeTLInputReportReasonGeoIrrelevant(nil).To_ReportReason()
	case REASON_FAKE:
		reason = mtproto.MakeTLInputReportReasonFake(nil).To_ReportReason()
	default:
		reason = mtproto.MakeTLInputReportReasonOther(nil).To_ReportReason()
	}
	return
}
