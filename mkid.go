package mkid

import (
	"os"
	"strconv"
	"time"

	"github.com/denisbrodbeck/machineid"
)

func Mkid(lastTime int64, seq int) (res uint64, ltime int64, sequence int) {
	const BASETIME int64 = 161401067
	const seqMask int = 4095
	id, _ := machineid.ID()
	mid, _ := strconv.ParseUint(id, 16, 0)
	t := time.Now().Unix()
	if t < BASETIME {
		panic("Time Error")
	}
	if lastTime == t {
		seq = (seq + 1) & seqMask
		if seq == 0 {
			time.Sleep(time.Millisecond)
		}
	} else {
		seq = 0
	}
	res = uint64(((t-BASETIME)<<22)|(((int64(os.Getpid())+int64(mid))<<12)&4194303)|int64(seq)) & 9223372036854775807
	return res, t, seq
}
