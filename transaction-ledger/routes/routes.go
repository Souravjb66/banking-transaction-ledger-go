package routes
import(
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"tr-ledger/service"
	// "tr-ledger/connect"
	"tr-ledger/config"

)
func Routes(){
	app:=fiber.New()
	app.Use(logger.New())
	app.Get("ac/call",service.Transaction)

	config.GetConnection().App=app





}