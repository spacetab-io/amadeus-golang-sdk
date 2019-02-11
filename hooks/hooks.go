package hooks

// A hook to be fired when logging on the logging levels returned from
// `Levels()` on your implementation of the interface. Note that this is not
// fired in a goroutine or a channel with workers, you should handle such
// functionality yourself if your call is non-blocking and you don't wish for
// the logging calls for levels returned from `Levels()` to block.

type Hook interface {
	Levels() []string
	Fire(string) error
}

// Internal type for storing the hooks on a logger instance.
type LevelHooks map[string][]Hook

// Add a hook to an instance of logger. This is called with
// `log.Hooks.Add(new(MyHook))` where `MyHook` implements the `Hook` interface.
func (hooks LevelHooks) Add(hook Hook) {
	if hooks == nil {
		hooks = make(LevelHooks)
	}

	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

// Fire all the hooks for the passed level. Used by `entry.log` to fire
// appropriate hooks for a log entry.
func (hooks LevelHooks) Fire(level string, entry string) error {
	for _, hook := range hooks[level] {
		if err := hook.Fire(entry); err != nil {
			return err
		}
	}

	return nil
}

//type Hooks []Hook
//
//type Hook struct {
//	hook func()
//	level string
//}
//
//func (h Hooks) Add(hook Hook) {
//	h = append(h, hook)
//}
//
//func (h Hooks) Log(level string, value string) {
//	for _, hook := range h {
//		if hook.level == level {
//			hook.hook()
//		}
//	}
//}

//var hooks map[string][]func(string)
//
//func Init() {
//	hooks = make(map[string][]func(string))
//}
//
//func Add(level string, hook func(string)) {
//	hooks[level] = append(hooks[level], hook)
//}
//
//func Fire(level string, entry string) {
//	for _, hook := range hooks[level] {
//		hook(entry)
//	}
//}
