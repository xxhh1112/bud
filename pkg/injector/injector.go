package injector

import (
	"log/slog"
	"os"

	"github.com/livebud/bud/pkg/bud"
	"github.com/livebud/bud/pkg/cli"
	"github.com/livebud/bud/pkg/controller"
	"github.com/livebud/bud/pkg/di"
	"github.com/livebud/bud/pkg/env"
	"github.com/livebud/bud/pkg/log"
	"github.com/livebud/bud/pkg/mod"
	"github.com/livebud/bud/pkg/mux"
	"github.com/livebud/bud/pkg/view"
	"github.com/livebud/bud/pkg/view/gohtml"
	"github.com/livebud/bud/pkg/web"
)

func New() di.Injector {
	in := di.New()
	di.Provide[*mod.Module](in, mod.Find)
	di.Provide[*cli.CLI](in, cli.Default)
	di.Provide[cli.Parser](in, cliParser)
	di.Provide[cli.Command](in, cliCommand)
	di.Provide[*bud.Command](in, bud.New)
	di.Register[*cli.CLI](in, bud.Register)
	di.Provide[*env.Bud](in, env.Load[*env.Bud])
	di.Provide[*slog.Logger](in, log.Default)
	di.Provide[*mux.Router](in, mux.New)
	di.Provide[*web.Server](in, webServer)
	di.Provide[web.Handler](in, webHandler)
	di.Provide[web.Router](in, webRouter)
	di.Provide[*controller.Router](in, controller.New)
	di.Provide[*view.Viewer](in, viewViewer)
	di.Provide[view.Finder](in, viewFinder)
	return in
}

func cliParser(cli *cli.CLI) cli.Parser {
	return cli
}

func cliCommand(cli *cli.CLI) cli.Command {
	return cli
}

func webHandler(router web.Router) web.Handler {
	return router
}

func webRouter(router *mux.Router) web.Router {
	return router
}

func webServer(budEnv *env.Bud, handler web.Handler) *web.Server {
	return &web.Server{
		Addr:    budEnv.Listen,
		Handler: handler,
	}
}

func viewViewer(mod *mod.Module) *view.Viewer {
	fsys := os.DirFS(mod.Directory("view"))
	return view.New(fsys, map[string]view.Renderer{
		".gohtml": gohtml.New(),
	})
}

func viewFinder(viewer *view.Viewer) view.Finder {
	return viewer
}
