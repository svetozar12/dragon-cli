package installers

var DependencyVersionMap = map[string]string{
	// Nx
	"nx":                "16.7.4",
	"@nx/workspace":     "16.7.4",
	"@nx/cypress":       "^16.7.4",
	"@nx/esbuild":       "^16.7.4",
	"@nx/eslint-plugin": "^16.7.4",
	"@nx/express":       "^16.7.4",
	"@nx/jest":          "^16.7.4",
	"@nx/linter":        "^16.7.4",
	"@nx/node":          "^16.7.4",
	"@nx/react":         "^16.7.4",
	"@nx/vite":          "^16.7.4",
	"@nx/webpack":       "^16.7.4",
	// Axios
	"axios": "^1.5.0",
	// Swagger
	"swagger-jsdoc":                       "^6.2.8",
	"swagger-ui-express":                  "^5.0.0",
	"@openapitools/openapi-generator-cli": "^2.7.0",
	// Typescript
	"tslib":                            "^2.6.2",
	"@types/express":                   "^4.17.17",
	"@types/jest":                      "^29.5.4",
	"@types/node":                      "^20.5.9",
	"@types/react":                     "^18.2.21",
	"@types/react-dom":                 "^18.2.7",
	"@types/swagger-jsdoc":             "^6.0.1",
	"@types/swagger-ui-express":        "^4.1.3",
	"@typescript-eslint/eslint-plugin": "^6.5.0",
	"@typescript-eslint/parser":        "^6.5.0",
	"ts-jest":                          "^29.1.1",
	"ts-node":                          "^10.9.1",
	"zod":                              "^3.22.2",
	// Swc
	"@swc/cli":     "^0.1.62",
	"@swc/core":    "^1.3.82",
	"@swc/helpers": "^0.5.1",
	// vite
	"@vitejs/plugin-react": "^4.0.4",
	"@vitest/coverage-c8":  "^0.33.0",
	"@vitest/ui":           "^0.34.3",
	"vite":                 "^4.4.9",
	"vitest":               "^0.34.3",
	// eslint
	"eslint":                    "^8.48.0",
	"eslint-config-prettier":    "^9.0.0",
	"eslint-plugin-cypress":     "^2.14.0",
	"eslint-plugin-import":      "^2.28.1",
	"eslint-plugin-jsx-a11y":    "^6.7.1",
	"eslint-plugin-react":       "^7.33.2",
	"eslint-plugin-react-hooks": "^4.6.0",
	// react
	"react":            "^18.2.0",
	"react-dom":        "^18.2.0",
	"react-router-dom": "^6.15.0",
	// jest
	"jest":                  "^29.6.4",
	"jest-environment-node": "^29.6.4",
	"jsdom":                 "^22.1.0",
	"cypress":               "^13.1.0",
	// astro
	"astro": "^3.0.7",
	// mongoose
	"mongoose":        "^7.5.0",
	"@types/mongoose": "^5.11.97",
}

type AvailableDependencies string

const (
	NextAuth                  = "next-auth"
	NextAuthPrisma            = "@next-auth/prisma-adapter"
	Prisma                    = "prisma"
	PrismaClient              = "@prisma/client"
	TailwindCSS               = "tailwindcss"
	Autoprefixer              = "autoprefixer"
	PostCSS                   = "postcss"
	Prettier                  = "prettier"
	PrettierPluginTailwindCSS = "prettier-plugin-tailwindcss"
	TRPCClient                = "@trpc/client"
	TRPCServer                = "@trpc/server"
	TRPCReactQuery            = "@trpc/react-query"
	TRPCNext                  = "@trpc/next"
	TanstackReactQuery        = "@tanstack/react-query"
	SuperJSON                 = "superjson"
)
