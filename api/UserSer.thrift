namespace go User

typedef i64 Timestamp

struct User {
  1:  required  string id,
  2:  required  string username,
  3:  required  string email,
  4:  required  string name,
  6:  optional  string timezone,
  9:  optional  Timestamp created,
  10: optional  Timestamp updated,
  11: optional  Timestamp deleted,
  13: optional  bool active,
  14: optional  string shardId,
  15: optional  UserAttributes attributes,
}

struct UserAttributes {
  1:  optional  string defaultLocationName,
  2:  optional  double defaultLatitude,
  3:  optional  double defaultLongitude,
  4:  optional  bool preactivation,
  5:  optional  list<string> viewedPromotions,
  6:  optional  string incomingEmailAddress,
  7:  optional  list<string> recentMailedAddresses,
  9:  optional  string comments,
  11: optional  Timestamp dateAgreedToTermsOfService,
  12: optional  i32 maxReferrals,
  13: optional  i32 referralCount,
  14: optional  string refererCode,
  15: optional  Timestamp sentEmailDate,
  16: optional  i32 sentEmailCount,
  17: optional  i32 dailyEmailLimit,
  18: optional  Timestamp emailOptOutDate,
  19: optional  Timestamp partnerEmailOptInDate,
  20: optional  string preferredLanguage,
  21: optional  string preferredCountry,
  22: optional  bool clipFullPage,
  23: optional  string twitterUserName,
  24: optional  string twitterId,
  25: optional  string groupName,
  26: optional  string recognitionLanguage,
  28: optional  string referralProof,
  29: optional  bool educationalDiscount,
  30: optional  string businessAddress,
  31: optional  bool hideSponsorBilling,
  32: optional  bool taxExempt,
  33: optional  bool useEmailAutoFiling,
  34: optional  i32 reminderEmailConfig
}

service UserSer {
  string Register(1:string u)
  string Login(1:string key, 2:string pwd)
  string Logout(1:string key)
}

