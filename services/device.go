package services

import (
	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/ElioenaiFerrari/security-dog-api/views"
	"gorm.io/gorm"
)

type DeviceService struct {
	db *gorm.DB
}

func NewDeviceService(db *gorm.DB) *DeviceService {
	return &DeviceService{db: db}
}

func (deviceService *DeviceService) Add(userID, remoteIP string) (views.DeviceView, error) {
	var device entities.Device
	var deviceView views.DeviceView

	if err := deviceService.db.Where("user_id = ? AND remote_ip = ?", userID, remoteIP).First(&device).Scan(&deviceView).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return deviceView, err
		} else {
			device.UserID = userID
			device.RemoteIP = remoteIP
			device.IsLinked = true

			if err := deviceService.db.Create(&device).Scan(&deviceView).Error; err != nil {
				return deviceView, err
			}

			return deviceView, nil

		}
	} else {
		device.IsLinked = true

		if err := deviceService.db.Updates(&device).Scan(&deviceView).Error; err != nil {
			return deviceView, err
		}

		return deviceView, nil
	}
}

func (deviceService *DeviceService) GetAll(userID string) ([]views.DeviceView, error) {
	var devicesView []views.DeviceView

	if err := deviceService.db.Where("user_id = ?", userID).Find(&entities.Device{}).Scan(&devicesView).Error; err != nil {
		return devicesView, err
	}

	return devicesView, nil
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
