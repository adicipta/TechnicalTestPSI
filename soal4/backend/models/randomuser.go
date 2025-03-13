package models

type RandomUserResponse struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			Street struct {
				Number int    `json:"number"`
				Name   string `json:"name"`
			} `json:"street"`
			City    string `json:"city"`
			State   string `json:"state"`
			Country string `json:"country"`
		} `json:"location"`
		Email string `json:"email"`
		Dob   struct {
			Age int `json:"age"`
		} `json:"dob"`
		Phone   string `json:"phone"`
		Cell    string `json:"cell"`
		Picture struct {
			Large  string `json:"large"`
			Medium string `json:"medium"`
			Thumb  string `json:"thumbnail"`
		} `json:"picture"`
	} `json:"results"`
}
