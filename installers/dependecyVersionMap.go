package installers

var DependencyVersionMap = map[string]string{
	// Nx
	NX:               "16.7.4",
	NX_GO:            "^2.7.0",
	NX_WORKSPACE:     "16.7.4",
	NX_CYPRESS:       "16.7.4",
	NX_ESBUILD:       "16.7.4",
	NX_ESLINT_PLUGIN: "16.7.4",
	NX_EXPRESS:       "16.7.4",
	NX_JEST:          "16.7.4",
	NX_LINTER:        "16.7.4",
	NX_NODE:          "16.7.4",
	NX_REACT:         "16.7.4",
	NX_VITE:          "16.7.4",
	NX_WEBPACK:       "16.7.4",
	NX_NEXT:          "16.7.4",
	// Axios
	AXIOS: "^1.5.0",
	// Swagger
	SWAGGER_JSDOC:            "^6.2.8",
	SWAGGER_UI_EXPRESS:       "^5.0.0",
	OPEN_API_GENERATOR_CLI:   "^2.7.0",
	TYPES_SWAGGER_JSDOC:      "^6.0.1",
	TYPES_SWAGGER_UI_EXPRESS: "^4.1.3",
	// Typescript
	TYPESCRIPT:               "^5.2.2",
	TSLIB:                    "^2.6.2",
	TYPES_EXPRESS:            "^4.17.17",
	TYPES_NODE:               "^20.5.9",
	TSC_ESLINT_ESLINT_PLUGIN: "^6.5.0",
	TSC_ESLINT_PARSER:        "^6.5.0",
	TS_NODE:                  "^10.9.1",
	ZOD:                      "^3.22.2",
	// Swc
	SWC_CLI:     "^0.1.62",
	SWC_CORE:    "^1.3.82",
	SWC_HELPERS: "^0.5.1",
	// vite
	VITEST_PLUGIN_REACT: "^4.0.4",
	VITEST_COVERAGE:     "^0.33.0",
	VITEST_UI:           "^0.34.3",
	VITE:                "^4.4.9",
	VITEST:              "^0.34.3",
	// eslint
	ESLINT:                    "^8.48.0",
	ESLINT_CONFIG_PRETTIER:    "^9.0.0",
	ESLINT_PLUGIN_CYPRESS:     "^2.14.0",
	ESLINT_PLUGIN_IMPORT:      "^2.28.1",
	ESLINT_PLUGIN_JSX_A11Y:    "^6.7.1",
	ESLINT_PLUGIN_REACT:       "^7.33.2",
	ESLINT_PLUGIN_REACT_HOOKS: "^4.6.0",
	// react
	NEXT:             "^13.1.12",
	REACT:            "^18.2.0",
	REACT_DOM:        "^18.2.0",
	REACT_ROUTER_DOM: "^6.15.0",
	TYPES_REACT:      "^18.2.21",
	TYPES_REACT_DOM:  "^18.2.7",
	// jest
	JEST:          "^29.6.4",
	JEST_ENV_NODE: "^29.6.4",
	JSDOM:         "^22.1.0",
	CYPRESS:       "^13.1.0",
	TYPES_JEST:    "^29.5.4",
	TS_JEST:       "^29.1.1",
	// astro
	ASTRO: "^3.0.7",
	// mongoose
	MONGOOSE:       "^7.5.0",
	TYPES_MONGOOSE: "^5.11.97",
}

type AvailableDependencies string

const (
	// VITE
	VITEST_PLUGIN_REACT = "@vitejs/plugin-react"
	VITEST_COVERAGE     = "@vitest/coverage-c8"
	VITEST_UI           = "@vitest/ui"
	VITE                = "vite"
	VITEST              = "vitest"
	// SWC
	SWC_CLI     = "@swc/cli"
	SWC_CORE    = "@swc/core"
	SWC_HELPERS = "@swc/helpers"
	// AXIOS
	AXIOS = "axios"
	// SWAGGER
	SWAGGER_JSDOC            = "swagger-jsdoc"
	SWAGGER_UI_EXPRESS       = "swagger-ui-express"
	TYPES_SWAGGER_JSDOC      = "@types/swagger-jsdoc"
	TYPES_SWAGGER_UI_EXPRESS = "@types/swagger-ui-express"
	OPEN_API_GENERATOR_CLI   = "@openapitools/openapi-generator-cli"
	// ESLINT
	ESLINT                    = "eslint"
	ESLINT_CONFIG_PRETTIER    = "eslint-config-prettier"
	ESLINT_PLUGIN_CYPRESS     = "eslint-plugin-cypress"
	ESLINT_PLUGIN_IMPORT      = "eslint-plugin-import"
	ESLINT_PLUGIN_JSX_A11Y    = "eslint-plugin-jsx-a11y"
	ESLINT_PLUGIN_REACT       = "eslint-plugin-react"
	ESLINT_PLUGIN_REACT_HOOKS = "eslint-plugin-react-hooks"
	// REACT
	REACT            = "react"
	REACT_DOM        = "react-dom"
	REACT_ROUTER_DOM = "react-router-dom"
	TYPES_REACT      = "@types/react"
	TYPES_REACT_DOM  = "@types/react-dom"
	NEXT             = "next"

	// NX
	NX               = "nx"
	NX_GO            = "@nx-go/nx-go"
	NX_CYPRESS       = "@nx/cypress"
	NX_ESBUILD       = "@nx/esbuild"
	NX_ESLINT_PLUGIN = "@nx/eslint-plugin"
	NX_EXPRESS       = "@nx/express"
	NX_JEST          = "@nx/jest"
	NX_LINTER        = "@nx/linter"
	NX_NODE          = "@nx/node"
	NX_REACT         = "@nx/react"
	NX_VITE          = "@nx/vite"
	NX_WEBPACK       = "@nx/webpack"
	NX_WORKSPACE     = "@nx/workspace"
	NX_NEXT          = "@nx/next"
	// TYPESCRIPT
	TYPESCRIPT               = "typescript"
	TSC_ESLINT_ESLINT_PLUGIN = "@typescript-eslint/eslint-plugin"
	TSC_ESLINT_PARSER        = "@typescript-eslint/parser"
	TS_NODE                  = "ts-node"
	ZOD                      = "zod"
	TYPES_NODE               = "@types/node"
	TSLIB                    = "tslib"
	TYPES_EXPRESS            = "@types/express"
	// JEST
	JEST          = "jest"
	JEST_ENV_NODE = "jest-environment-node"
	JSDOM         = "jsdom"
	CYPRESS       = "cypress"
	TYPES_JEST    = "@types/jest"
	TS_JEST       = "ts-jest"
	// astro
	ASTRO = "astro"
	// mongoose
	MONGOOSE       = "mongoose"
	TYPES_MONGOOSE = "@types/mongoose"
)
