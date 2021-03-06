// Tonika: A distributed social networking platform
// Copyright (C) 2010 Petar Maymounkov <petar@5ttt.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.


package sys

import (
	"os"
)

type Bank interface {
	String() string
	GetBuild() string

	GetMyId() Id
	GetMyName() string
	GetMyEmail() string
	GetMyAddr() string
	GetMyExtAddr() string
	GetMySignatureKey() *SigPubKey

	SetMy(key string, v interface{})

	GetBySlot(slot int) (View, os.Error)
	GetById(id Id) (View, os.Error)
	GetByAcceptKey(ak *DialKey) (View, os.Error)
	GetByDialKey(dk *DialKey) (View, os.Error)

	Reserve() View
	Revoke(slot int)
	Enumerate() []View
	Write(slot int, key string, v interface{}) (View, os.Error)
	Sync(slot int)
	SyncAddr(slot int)
	Save()
}
