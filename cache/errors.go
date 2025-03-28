package cache

import "errors"

var (
	ErrInvalidOp       = errors.New("invalid op")
	ErrShmNotInit      = errors.New("shm not init")
	ErrShmAlreadyInit  = errors.New("shm already init")
	ErrShmVersion      = errors.New("invalid shm version")
	ErrShmSize         = errors.New("invalid shm size")
	ErrShmNotLoaded    = errors.New("shm not loaded")
	ErrAddToUHash      = errors.New("unable to add-to-uhash")
	ErrRemoveFromUHash = errors.New("unable to remove-from-uhash")

	ErrInvalidUID = errors.New("invalid uid")

	ErrStats = errors.New("invalid stats")

	ErrNotFound = errors.New("not found")
	ErrBusy     = errors.New("busy")

	ErrInvalidNumBoards = errors.New("invalid num boards")

	ErrMapNotInit     = errors.New("map not init")
	ErrMapAlreadyInit = errors.New("map already init")
)
