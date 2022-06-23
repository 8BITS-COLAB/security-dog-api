package services

import (
	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"gorm.io/gorm"
)

type DeviceService struct {
	db *gorm.DB
}

func NewDeviceService(db *gorm.DB) *DeviceService {
	return &DeviceService{db: db}
}

func (deviceService *DeviceService) Add(userID, remoteIP string) (entities.Device, error) {

	var device entities.Device

	if err := deviceService.db.Where("user_id = ? AND remote_ip = ?", userID, remoteIP).First(&device).Error; err != nil {
		return device, err
	}

	if device.ID != "" {
		device.IsLinked = true

		if err := deviceService.db.Updates(&device).Error; err != nil {
			return device, err
		}

		return device, nil
	} else {
		if err := deviceService.db.Create(&device).Error; err != nil {
			return device, err
		}

		return device, nil
	}
}

func (deviceService *DeviceService) GetAll(userID string) ([]entities.Device, error) {
	var devices []entities.Device

	if err := deviceService.db.Where("user_id = ?", userID).Find(&devices).Error; err != nil {
		return devices, err
	}

	return devices, nil
}

func (deviceService *DeviceService) Update(updateDeviceDTO *dtos.UpdateDeviceDTO) (entities.Device, error) {
	device := entities.Device{}

	if err := deviceService.db.Where("user_id = ? AND remote_ip = ?", updateDeviceDTO.UserID, updateDeviceDTO.RemoteIP).First(&device).Error; err != nil {
		return device, err
	}

	device.IsLinked = updateDeviceDTO.IsLinked
	device.IsTrusted = updateDeviceDTO.IsTrusted
	device.IsBlocked = updateDeviceDTO.IsBlocked

	if err := deviceService.db.Updates(&device).Error; err != nil {
		return device, err
	}

	return device, nil
}
