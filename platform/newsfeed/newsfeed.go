package newsfeed

type NewsFeedRepo struct {
	ID          uint64 `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

/**
 * TODO: specify table name
 *
 * @return string
 */
func (nfr *NewsFeedRepo) TableName() string {
	return "newsfeed"
}

/**
 * TODO: update specify data
 *
 * @param "data" map[string]interface{} `NewsFeedRepo convert to map[]`
 *
 * @return err
 */
func (nfr *NewsFeedRepo) Update(data map[string]interface{}) (err error) {
	return nil
}

/**
 * TODO: get data in list
 *
 * @param "id" string
 *
 * @return err
 */
func (nfr *NewsFeedRepo) Get(id string) (err error) {
	return nil
}

/**
 * TODO: store data
 *
 * @return err
 */
func (nfr *NewsFeedRepo) Store() (err error) {
	return nil
}

/***
 * TODO: find specify data
 *
 * @param "id" string
 *
 * @return err
 */
func (nfr *NewsFeedRepo) FindId(id string) (err error) {
	return nil
}
