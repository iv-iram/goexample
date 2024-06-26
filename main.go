package main

import (
	"fmt"
	"github.com/iv-iram/goexamples/hello-world"
	"github.com/iv-iram/goexamples/values"
	"github.com/iv-iram/goexamples/variables"
	"github.com/iv-iram/goexamples/constants"
	forloop "github.com/iv-iram/goexamples/for" // Use alias for package named `for`
	ifelse "github.com/iv-iram/goexamples/if-else" // Use alias for package named `if-else`
	switchcase "github.com/iv-iram/goexamples/switch" // Use alias for package named `switch`
	"github.com/iv-iram/goexamples/arrays"
	"github.com/iv-iram/goexamples/slices"
	"github.com/iv-iram/goexamples/maps"
	"github.com/iv-iram/goexamples/range"
	"github.com/iv-iram/goexamples/functions"
	multipleReturns "github.com/iv-iram/goexamples/multiple-return-values" // Use alias for package named `multiple-return-values`
	"github.com/iv-iram/goexamples/variadic-functions"
	"github.com/iv-iram/goexamples/closures"
	"github.com/iv-iram/goexamples/recursions"
	"github.com/iv-iram/goexamples/pointer"
	"github.com/iv-iram/goexamples/string-runes"
	"github.com/iv-iram/goexamples/structs"
	"github.com/iv-iram/goexamples/methods"
	"github.com/iv-iram/goexamples/interfaces"
	"github.com/iv-iram/goexamples/enums"
	"github.com/iv-iram/goexamples/struct-embeddings"
	"github.com/iv-iram/goexamples/generics"
	"github.com/iv-iram/goexamples/errors"
	"github.com/iv-iram/goexamples/custom-errors"
	"github.com/iv-iram/goexamples/goroutines"
	"github.com/iv-iram/goexamples/channels"
	channelBuffering "github.com/iv-iram/goexamples/channel-buffering" // Use alias for package named `channel-buffering`
	channelSync "github.com/iv-iram/goexamples/channel-synchronisation" // Use alias for package named `channel-synchronisation`
	channelDirections "github.com/iv-iram/goexamples/channel-directions" // Use alias for package named `channel-directions`
	"github.com/iv-iram/goexamples/select"
	"github.com/iv-iram/goexamples/timeouts"
	nonBlocking "github.com/iv-iram/goexamples/non-blocking-channel-operations" // Use alias for package named `non-blocking-channel-operations`
	"github.com/iv-iram/goexamples/closing-channels"
	"github.com/iv-iram/goexamples/range-over-channels"
	"github.com/iv-iram/goexamples/timers"
	"github.com/iv-iram/goexamples/tickers"
	"github.com/iv-iram/goexamples/worker-pools"
	"github.com/iv-iram/goexamples/wait-group"
	"github.com/iv-iram/goexamples/rate-limiting"
	"github.com/iv-iram/goexamples/atomic-counters"
	"github.com/iv-iram/goexamples/mutexes"
	"github.com/iv-iram/goexamples/defer"
)

func main() {
	fmt.Println("Running hello-world example:")
	hello_world.Hello_world()
	fmt.Println("\n")

	fmt.Println("Running values example:")
	values.Values()
	fmt.Println("\n")

	fmt.Println("Running variables example:")
	variables.Variables()
	fmt.Println("\n")

	fmt.Println("Running constants example:")
	constants.Constants()
	fmt.Println("\n")

	fmt.Println("Running for example:")
	forloop.Forloop()
	fmt.Println("\n")

	fmt.Println("Running if-else example:")
	ifelse.If_else()
	fmt.Println("\n")

	fmt.Println("Running switch example:")
	switchcase.Switch_go()
	fmt.Println("\n")

	fmt.Println("Running arrays example:")
	arrays.Arrays()
	fmt.Println("\n")

	fmt.Println("Running slices example:")
	slice.Slice()
	fmt.Println("\n")

	fmt.Println("Running maps example:")
	maps.Maps()
	fmt.Println("\n")

	fmt.Println("Running range example:")
	ranges.Ranges()
	fmt.Println("\n")

	fmt.Println("Running functions example:")
	functions.Functions()
	fmt.Println("\n")

	fmt.Println("Running multiple-return-values example:")
	multipleReturns.Multiple_return_functions()
	fmt.Println("\n")

	fmt.Println("Running variadic-functions example:")
	variadic_functions.Variadic_functions()
	fmt.Println("\n")

	fmt.Println("Running closures example:")
	closures.Closures()
	fmt.Println("\n")

	fmt.Println("Running recursions example:")
	recursions.Recursions()
	fmt.Println("\n")

	fmt.Println("Running pointer example:")
	pointer.Pointer()
	fmt.Println("\n")

	fmt.Println("Running string-runes example:")
	string_runes.String_runes()
	fmt.Println("\n")

	fmt.Println("Running structs example:")
	struct_go.Struct_go()
	fmt.Println("\n")

	fmt.Println("Running methods example:")
	methods.Methods()
	fmt.Println("\n")

	fmt.Println("Running interfaces example:")
	interface_go.Interface_go()
	fmt.Println("\n")

	fmt.Println("Running enums example:")
	enums.Enum()
	fmt.Println("\n")

	fmt.Println("Running struct-embeddings example:")
	struct_embeddings.Struct_embeddings()
	fmt.Println("\n")

	fmt.Println("Running generics example:")
	generics.Generics()
	fmt.Println("\n")

	fmt.Println("Running errors example:")
	errors.Errors()
	fmt.Println("\n")

	fmt.Println("Running custom-errors example:")
	custom_errors.Custom_errors()
	fmt.Println("\n")

	fmt.Println("Running goroutines example:")
	goroutines.Goroutines()
	fmt.Println("\n")

	fmt.Println("Running channels example:")
	channels.Channels()
	fmt.Println("\n")

	fmt.Println("Running channel-buffering example:")
	channelBuffering.Channel_bufferings()
	fmt.Println("\n")

	fmt.Println("Running channel-synchronisation example:")
	channelSync.Channel_synchronisations()
	fmt.Println("\n")

	fmt.Println("Running channel-directions example:")
	channelDirections.Channel_directions()
	fmt.Println("\n")

	fmt.Println("Running select example:")
	select_go.Select_go()
	fmt.Println("\n")

	fmt.Println("Running timeouts example:")
	timeouts.Timeouts()
	fmt.Println("\n")

	fmt.Println("Running non-blocking-channel-operations example:")
	nonBlocking.Non_blocking_channel()
	fmt.Println("\n")

	fmt.Println("Running closing-channels example:")
	closing_channels.Closing_channels()
	fmt.Println("\n")

	fmt.Println("Running range-over-channels example:")
	range_over_channels.Range_over_channels()
	fmt.Println("\n")

	fmt.Println("Running timers example:")
	timers.Timers()
	fmt.Println("\n")

	fmt.Println("Running tickers example:")
	tickers.Tickers()
	fmt.Println("\n")

	fmt.Println("Running worker-pools example:")
	worker_pools.Worker_pools()
	fmt.Println("\n")

	fmt.Println("Running wait-group example:")
	wait_group.Wait_group()
	fmt.Println("\n")

	fmt.Println("Running rate-limiting example:")
	rate_limiting.Rate_limiting()
	fmt.Println("\n")

	fmt.Println("Running atomic-counters example:")
	atomic_counters.Atomic_counters()
	fmt.Println("\n")

	fmt.Println("Running mutexes example:")
	mutex_go.Mutex_go()
	fmt.Println("\n")

	fmt.Println("Running defer example")
	defer_go.Defer_go()
	fmt.Println("\n")

}
