package respdata

type BilibiliLoginInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Url          string `json:"url"`
		RefreshToken string `json:"refresh_token"`
		Timestamp    int    `json:"timestamp"`
		Code         int    `json:"code"`
		Message      string `json:"message"`
	} `json:"data"`
}

type BilibiliQRCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Url       string `json:"url"`
		QrcodeKey string `json:"qrcode_key"`
	} `json:"data"`
}

//type BilibiliLoginInfo struct {
//	Code    int    `json:"code"`
//	Message string `json:"message"`
//	Data    struct {
//		Url string `json:"url"`
//	} `json:"data"`
//}

type BilibiliCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
}

type BilibiliChouJiangNumData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Times int `json:"times"`
	} `json:"data"`
}

type BilibiliCaptcha struct {
	Data struct {
		Result struct {
			Gt         string `json:"gt"`
			Challenger string `json:"challenger"`
			Key        string `json:"key"`
		} `json:"result"`
	} `json:"data"`
}
type BilibiliHash struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type BilibiliSpaceHistory struct {
	Data struct {
		Cards []struct {
			Card string `json:"card"`
			Desc struct {
				OrigDyIdStr string `json:"orig_dy_id_str"` //父内容
				Origin      struct {
					Bvid         string `json:"bvid"`
					DynamicIdStr string `json:"dynamic_id_str"`
					RidStr       string `json:"rid_str"` //目标评论区id
					Type         int    `json:"type"`
					Uid          int    `json:"uid"` //对方uid
				} `json:"origin"`
				PreDyIdStr string `json:"pre_dy_id_str"` //本体转发文字动态id
				Previous   struct {
					DynamicIdStr string `json:"dynamic_id_str"` //子动态id
					RidStr       string `json:"rid_str"`        //子评论区id
					Type         int    `json:"type"`           //子动态类型
					Uid          int    `json:"uid"`            //对方uid
				} `json:"previous"`
			} `json:"desc"`
			ExtendJson string `json:"extend_json"`
			Display    struct {
				TopicInfo struct {
					NewTopic struct {
						Id   int    `json:"id"`
						Name string `json:"name"`
					} `json:"new_topic"`
				} `json:"topic_info"`
			} `json:"display"`
		} `json:"cards"`
		NextOffset int `json:"next_offset"`
	} `json:"data"`
	JsonData string
}

type BilibiliSpaceDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Item struct {
			Basic struct {
				CommentIdStr string `json:"comment_id_str"`
				CommentType  int    `json:"comment_type"`
				LikeIcon     struct {
					ActionUrl string `json:"action_url"`
					EndUrl    string `json:"end_url"`
					Id        int    `json:"id"`
					StartUrl  string `json:"start_url"`
				} `json:"like_icon"`
				RidStr string `json:"rid_str"`
			} `json:"basic"`
			IdStr   string `json:"id_str"`
			Modules struct {
				ModuleAuthor struct {
					Face           string `json:"face"`
					FaceNft        bool   `json:"face_nft"`
					Following      bool   `json:"following"`
					JumpUrl        string `json:"jump_url"`
					Label          string `json:"label"`
					Mid            int    `json:"mid"`
					Name           string `json:"name"`
					OfficialVerify struct {
						Desc string `json:"desc"`
						Type int    `json:"type"`
					} `json:"official_verify"`
					Pendant struct {
						Expire            int    `json:"expire"`
						Image             string `json:"image"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
						Name              string `json:"name"`
						Pid               int    `json:"pid"`
					} `json:"pendant"`
					PubAction       string `json:"pub_action"`
					PubLocationText string `json:"pub_location_text"`
					PubTime         string `json:"pub_time"`
					PubTs           int    `json:"pub_ts"`
					Type            string `json:"type"`
					Vip             struct {
						AvatarSubscript    int    `json:"avatar_subscript"`
						AvatarSubscriptUrl string `json:"avatar_subscript_url"`
						DueDate            int64  `json:"due_date"`
						Label              struct {
							BgColor               string `json:"bg_color"`
							BgStyle               int    `json:"bg_style"`
							BorderColor           string `json:"border_color"`
							ImgLabelUriHans       string `json:"img_label_uri_hans"`
							ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
							ImgLabelUriHant       string `json:"img_label_uri_hant"`
							ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
							LabelTheme            string `json:"label_theme"`
							Path                  string `json:"path"`
							Text                  string `json:"text"`
							TextColor             string `json:"text_color"`
							UseImgLabel           bool   `json:"use_img_label"`
						} `json:"label"`
						NicknameColor string `json:"nickname_color"`
						Status        int    `json:"status"`
						ThemeType     int    `json:"theme_type"`
						Type          int    `json:"type"`
					} `json:"vip"`
				} `json:"module_author"`
				ModuleDynamic struct {
					Additional interface{} `json:"additional"`
					Desc       struct {
						RichTextNodes []struct {
							OrigText string `json:"orig_text"`
							Text     string `json:"text"`
							Type     string `json:"type"`
							Rid      string `json:"rid,omitempty"`
						} `json:"rich_text_nodes"`
						Text string `json:"text"`
					} `json:"desc"`
					Major interface{} `json:"major"`
					Topic interface{} `json:"topic"`
				} `json:"module_dynamic"`
				ModuleMore struct {
					ThreePointItems []struct {
						Label string `json:"label"`
						Type  string `json:"type"`
					} `json:"three_point_items"`
				} `json:"module_more"`
				ModuleStat struct {
					Comment struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
					} `json:"comment"`
					Forward struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
					} `json:"forward"`
					Like struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
						Status    bool `json:"status"`
					} `json:"like"`
				} `json:"module_stat"`
			} `json:"modules"`
			Orig struct {
				Basic struct {
					CommentIdStr string `json:"comment_id_str"`
					CommentType  int    `json:"comment_type"`
					LikeIcon     struct {
						ActionUrl string `json:"action_url"`
						EndUrl    string `json:"end_url"`
						Id        int    `json:"id"`
						StartUrl  string `json:"start_url"`
					} `json:"like_icon"`
					RidStr string `json:"rid_str"`
				} `json:"basic"`
				IdStr   string `json:"id_str"`
				Modules struct {
					ModuleAuthor struct {
						Face           string      `json:"face"`
						FaceNft        bool        `json:"face_nft"`
						Following      interface{} `json:"following"`
						JumpUrl        string      `json:"jump_url"`
						Label          string      `json:"label"`
						Mid            int         `json:"mid"`
						Name           string      `json:"name"`
						OfficialVerify struct {
							Desc string `json:"desc"`
							Type int    `json:"type"`
						} `json:"official_verify"`
						Pendant struct {
							Expire            int    `json:"expire"`
							Image             string `json:"image"`
							ImageEnhance      string `json:"image_enhance"`
							ImageEnhanceFrame string `json:"image_enhance_frame"`
							Name              string `json:"name"`
							Pid               int    `json:"pid"`
						} `json:"pendant"`
						PubAction string `json:"pub_action"`
						PubTime   string `json:"pub_time"`
						PubTs     int    `json:"pub_ts"`
						Type      string `json:"type"`
						Vip       struct {
							AvatarSubscript    int    `json:"avatar_subscript"`
							AvatarSubscriptUrl string `json:"avatar_subscript_url"`
							DueDate            int64  `json:"due_date"`
							Label              struct {
								BgColor               string `json:"bg_color"`
								BgStyle               int    `json:"bg_style"`
								BorderColor           string `json:"border_color"`
								ImgLabelUriHans       string `json:"img_label_uri_hans"`
								ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
								ImgLabelUriHant       string `json:"img_label_uri_hant"`
								ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
								LabelTheme            string `json:"label_theme"`
								Path                  string `json:"path"`
								Text                  string `json:"text"`
								TextColor             string `json:"text_color"`
								UseImgLabel           bool   `json:"use_img_label"`
							} `json:"label"`
							NicknameColor string `json:"nickname_color"`
							Status        int    `json:"status"`
							ThemeType     int    `json:"theme_type"`
							Type          int    `json:"type"`
						} `json:"vip"`
					} `json:"module_author"`
					ModuleDynamic struct {
						Additional interface{} `json:"additional"`
						Desc       struct {
							RichTextNodes []RichTextNodes `json:"rich_text_nodes"`
							Text          string          `json:"text"`
						} `json:"desc"`
						Major struct {
							Draw struct {
								Id    int `json:"id"`
								Items []struct {
									Height int           `json:"height"`
									Size   float64       `json:"size"`
									Src    string        `json:"src"`
									Tags   []interface{} `json:"tags"`
									Width  int           `json:"width"`
								} `json:"items"`
							} `json:"draw"`
							Type string `json:"type"`
						} `json:"major"`
						Topic interface{} `json:"topic"`
					} `json:"module_dynamic"`
				} `json:"modules"`
				Type    string `json:"type"`
				Visible bool   `json:"visible"`
			} `json:"orig"`
			Type    string `json:"type"`
			Visible bool   `json:"visible"`
		} `json:"item"`
	} `json:"data"`
}

type BilibiliExtendJson struct {
	Data struct {
		Content string `json:"content"`
	} `json:"data"`
	Ctrl []struct {
		Data     string `json:"data"`
		Length   int    `json:"length"`
		Location int    `json:"location"`
		Type     int    `json:"type"`
	} `json:"ctrl"`
}

type BilibiliDynamicDetail struct {
	Code int `json:"code"`
	Data struct {
		Card struct {
			Card     string `json:"card"`
			CardJson struct {
				Item struct {
					Description string `json:"description"`
				} `json:"item"`
			}
		} `json:"card"`
	} `json:"data"`
}

type BilibiliActivity struct {
	BaseInfo struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		Keywords     string `json:"keywords"`
		SharePicture string `json:"sharePicture"`
	} `json:"BaseInfo"`
	LotteryNew []struct {
		LotteryId string `json:"lotteryId"`
	} `json:"h5-lottery-new"`
	FollowNew []struct {
		Uid string `json:"uid"`
	} `json:"h5-follow-new"`
	PcLotteryNew []struct {
		LotteryId string `json:"lotteryId"`
	} `json:"pc-lottery-new"`
	PcLotteryV3 []struct {
		LotteryId string `json:"lotteryId"`
	} `json:"pc-lottery-v3"`
	H5LotteryV3 []struct {
		LotteryId string `json:"lotteryId"`
	} `json:"h5-lottery-v3"`
}

type BilibiliSpaceHistoryCardJson struct {
	Item struct {
		Content string `json:"content"`
	}
	OriginUser struct {
		Info struct {
			Uid   int64  `json:"uid"`
			Uname string `json:"uname"`
		} `json:"info"`
	} `json:"origin_user"`
}

type BilibiliUserInfo struct {
	Mid   string `json:"mid"`
	Uname string `json:"uname"`
}

type BilibiliDo struct {
	Code int `json:"code"`
	Data []struct {
		GiftName string `json:"gift_name"`
	} `json:"data"`
	Message string `json:"message"`
}

type DynBody struct {
	DynReq struct {
		Content struct {
			Contents []DynContent `json:"contents"`
		} `json:"content"`
		Scene int `json:"scene"` //1创建动态，4转发
	} `json:"dyn_req"`
	WebRepostSrc struct {
		DynIdStr string `json:"dyn_id_str"`
	} `json:"web_repost_src"`
}

type DynContent struct {
	RawText string `json:"raw_text"` //文本
	Type    int    `json:"type"`     //1普通文本，2@，9表情
	BizId   string `json:"biz_id"`   //如果是@就需要填充对方uid
}

type RichTextNodes struct {
	JumpUrl  string `json:"jump_url,omitempty"`
	OrigText string `json:"orig_text"`
	Text     string `json:"text"`
	Type     string `json:"type"`
	Rid      string `json:"rid,omitempty"`
}

type Ctrl struct {
	Data     string `json:"data"`
	Length   int    `json:"length"`
	Location int    `json:"location"`
	Type     int    `json:"type"`
}

type ModifyList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		List []struct {
			Mid          int         `json:"mid"`
			Attribute    int         `json:"attribute"`
			Mtime        int         `json:"mtime"`
			Tag          interface{} `json:"tag"`
			Special      int         `json:"special"`
			ContractInfo struct {
				IsContractor bool `json:"is_contractor"`
				Ts           int  `json:"ts"`
				IsContract   bool `json:"is_contract"`
				UserAttr     int  `json:"user_attr"`
			} `json:"contract_info"`
			Uname          string `json:"uname"`
			Face           string `json:"face"`
			Sign           string `json:"sign"`
			FaceNft        int    `json:"face_nft"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Vip struct {
				VipType       int    `json:"vipType"`
				VipDueDate    int64  `json:"vipDueDate"`
				DueRemark     string `json:"dueRemark"`
				AccessStatus  int    `json:"accessStatus"`
				VipStatus     int    `json:"vipStatus"`
				VipStatusWarn string `json:"vipStatusWarn"`
				ThemeType     int    `json:"themeType"`
				Label         struct {
					Path        string `json:"path"`
					Text        string `json:"text"`
					LabelTheme  string `json:"label_theme"`
					TextColor   string `json:"text_color"`
					BgStyle     int    `json:"bg_style"`
					BgColor     string `json:"bg_color"`
					BorderColor string `json:"border_color"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				AvatarSubscriptUrl string `json:"avatar_subscript_url"`
			} `json:"vip"`
			NftIcon string `json:"nft_icon"`
		} `json:"list"`
		ReVersion int `json:"re_version"`
		Total     int `json:"total"`
	} `json:"data"`
}
