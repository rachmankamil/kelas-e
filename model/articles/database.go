package m_articles

import "km-kelas-e/config"

func SelectAll() (article []Article, err error) {
	if err = config.DB.Find(&article).Error; err != nil {
		return []Article{}, err
	}
	return
}
