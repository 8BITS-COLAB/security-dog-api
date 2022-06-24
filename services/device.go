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
			device.IsTrusted = true

			if err := deviceService.db.Create(&device).Scan(&deviceView).Error; err != nil {
				return deviceView, err
			}

			return deviceView, nil

		}
	} else {
		device.IsLinked = true
		device.IsTrusted = true

		if err := deviceService.db.UpdateColumns(&device).Scan(&deviceView).Error; err != nil {
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
	var device entities.Device

	if err := deviceService.db.Where("user_id = ? AND remote_ip = ?", updateDeviceDTO.UserID, updateDeviceDTO.RemoteIP).First(&device).Error; err != nil {
		return device, err
	}

	if err := deviceService.db.Model(&device).UpdateColumns(map[string]interface{}{
		"is_linked":  updateDeviceDTO.IsLinked,
		"is_trusted": updateDeviceDTO.IsTrusted,
		"is_blocked": updateDeviceDTO.IsBlocked,
	}).Error; err != nil {
		return device, err
	}

	return device, nil
}

func (deviceService *DeviceService) GetByRemoteIP(userID, remoteIP string) (entities.Device, error) {
	var device entities.Device

	if err := deviceService.db.Where("user_id = ? AND remote_ip = ?", userID, remoteIP).First(&device).Error; err != nil {
		return device, err
	}

	return device, nil
}
