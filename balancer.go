package cpubalancer

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/rurick/balancer/systemstat"
)

//Balancer - балансировщик нагрузки
type Balancer struct {
	value atomic.Value
}

//New - создать объект балансировщика и проинициализировать начальным значением
func New(val int) *Balancer {
	b := new(Balancer)
	b.value.Store(val)
	return b
}

//Run - запустить процесс слежения.
/**ctx - контекст выполнения
 * minCPUUsage - минимальное значение использование CPU  в % когда нужно инкрементировать значение val
 * maxCPUUsage - максимальное значение использование CPU  в % когда нужно декрементировать значение val
 * maxVal - максимально допустимое значение val
 * */
func (b *Balancer) Run(ctx context.Context, minCPUUsage int, maxCPUUsage int, maxVal int) {
	go func() {
		cpu1 := systemstat.GetCPUSample()
		cpu2 := cpu1
		ticker := time.NewTicker(1 * time.Second)

		for {
			select {
			case <-ticker.C:
				cpu1 = cpu2
				cpu2 = systemstat.GetCPUSample()
				st := systemstat.GetSimpleCPUAverage(cpu1, cpu2)
				val := b.value.Load().(int)
				if st.BusyPct > float64(maxCPUUsage) && val > 1 {
					val--
				}
				if st.BusyPct < float64(minCPUUsage) && val < maxVal {
					val++
				}
				b.value.Store(val)
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

//Value - получить значение
func (b *Balancer) Value() int {
	return b.value.Load().(int)
}

//Set - получить значение
func (b *Balancer) Set(val int) {
	b.value.Store(val)
}
