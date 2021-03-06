package migrations

import (
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"migrationScript/common"
	"migrationScript/models"
)

func Run() error {
	var modelsToMigrate []interface{}

	//TheAmazingCodeExample
	modelsToMigrate = append(modelsToMigrate, models.User{})
	modelsToMigrate = append(modelsToMigrate, models.ProfilePicture{})
	modelsToMigrate = append(modelsToMigrate, models.EmailConfirmation{})
	modelsToMigrate = append(modelsToMigrate, models.PhoneConfirmation{})
	modelsToMigrate = append(modelsToMigrate, models.Role{})
	modelsToMigrate = append(modelsToMigrate, models.PostalCode{})
	modelsToMigrate = append(modelsToMigrate, models.Permission{})
	modelsToMigrate = append(modelsToMigrate, models.Address{})

	//TheAmazingPostManager
	modelsToMigrate = append(modelsToMigrate, models.Post{})
	modelsToMigrate = append(modelsToMigrate, models.PostVote{})
	modelsToMigrate = append(modelsToMigrate, models.Comment{})
	modelsToMigrate = append(modelsToMigrate, models.CommentVote{})

	//TheAmazingNotificator
	modelsToMigrate = append(modelsToMigrate, models.Notification{})
	modelsToMigrate = append(modelsToMigrate, models.NotificationType{})

	r := common.GetDatabase()

	for _, model := range modelsToMigrate {

		r.DropTableIfExists(model)
		if r.Error != nil {
			return r.Error
		}

		r.AutoMigrate(model)
		if r.Error != nil {
			return r.Error
		}

	}

	err := createMockData()
	if err != nil {
		return err
	}

	err = createTriggers()
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func createMockData() error {
	println("Start migrations")
	var err error

	notiType1 := models.NotificationType{
		ID:   1,
		Type: "Post vote",
	}
	notiType2 := models.NotificationType{
		ID:   2,
		Type: "Post comment",
	}
	notiType3 := models.NotificationType{
		ID:   3,
		Type: "Comment comment",
	}
	err = common.GetDatabase().Create(&notiType1).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&notiType2).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&notiType3).Error
	if err != nil {
		return err
	}

	//Permissions
	permission1 := models.Permission{
		ID:          1,
		Description: "Profile",
	}
	permission2 := models.Permission{
		ID:          2,
		Description: "User Management",
	}
	permission3 := models.Permission{
		ID:          3,
		Description: "Role Management",
	}

	err = common.GetDatabase().Create(&permission1).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&permission2).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&permission3).Error
	if err != nil {
		return err
	}

	//Roles
	role1 := models.Role{
		ID:          1,
		Description: "Admin",
		Permissions: []models.Permission{permission1, permission2, permission3},
	}
	role2 := models.Role{
		ID:          2,
		Description: "User",
		Permissions: []models.Permission{permission1},
	}

	err = common.GetDatabase().Create(&role1).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&role2).Error
	if err != nil {
		return err
	}

	//Postal codes
	postalCode1 := models.PostalCode{
		ID:         1,
		PostalCode: "C1405",
	}
	postalCode2 := models.PostalCode{
		ID:         2,
		PostalCode: "C1424",
	}
	postalCode3 := models.PostalCode{
		ID:         3,
		PostalCode: "C1407",
	}
	postalCode4 := models.PostalCode{
		ID:         4,
		PostalCode: "C1424",
	}
	postalCode5 := models.PostalCode{
		ID:         5,
		PostalCode: "C1406",
	}
	postalCode6 := models.PostalCode{
		ID:         6,
		PostalCode: "1416",
	}

	err = common.GetDatabase().Create(&postalCode1).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&postalCode2).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&postalCode3).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&postalCode4).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&postalCode5).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&postalCode6).Error
	if err != nil {
		return err
	}

	//Users
	defaultPassword, _ := hashPassword("password")

	user1 := models.User{
		GUID:     xid.New().String(),
		Name:     "Juan Manuel",
		LastName: "Oubina",
		Password: defaultPassword,
		Email:    "juan.manuel.oubina@gmail.com",
		Phone:    "",
		RoleID:   1,
	}
	user1.ID = 1
	user2 := models.User{
		GUID:     xid.New().String(),
		Name:     "John",
		LastName: "Doe",
		Password: defaultPassword,
		Email:    "johnDoe@gmail.com",
		Phone:    "",
		RoleID:   2,
	}
	user1.ID = 2

	err = common.GetDatabase().Create(&user1).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&user2).Error
	if err != nil {
		return err
	}

	//Addresses
	address1 := models.Address{
		ID:           1,
		Address:      "Angel Gallardo 551",
		Floor:        "2",
		Apartment:    "A",
		MainAddress:  true,
		PostalCodeID: 1,
		UserID:       1,
	}
	address2 := models.Address{
		ID:           2,
		Address:      "San Jose de Calasanz 524",
		MainAddress:  false,
		PostalCodeID: 2,
		UserID:       1,
	}
	address3 := models.Address{
		ID:           3,
		Address:      "Santander 4000",
		MainAddress:  true,
		PostalCodeID: 6,
		UserID:       2,
	}

	err = common.GetDatabase().Create(&address1).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&address2).Error
	if err != nil {
		return err
	}
	err = common.GetDatabase().Create(&address3).Error
	if err != nil {
		return err
	}

	return nil

}

func createTriggers() error {

	data, err := ioutil.ReadFile("triggers.sql")
	if err != nil {
		return err
	}

	err = common.GetDatabase().Exec(string(data)).Error
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
