// Code generated from metric.g4 by ANTLR 4.7.2. DO NOT EDIT.

package metric // metric
import "github.com/antlr/antlr4/runtime/Go/antlr"

// metricListener is a complete listener for a parse tree produced by metricParser.
type metricListener interface {
	antlr.ParseTreeListener

	// EnterUom is called when entering the uom production.
	EnterUom(c *UomContext)

	// EnterMeasure is called when entering the measure production.
	EnterMeasure(c *MeasureContext)

	// EnterExponent is called when entering the exponent production.
	EnterExponent(c *ExponentContext)

	// EnterPrefix is called when entering the prefix production.
	EnterPrefix(c *PrefixContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterBaseunit is called when entering the baseunit production.
	EnterBaseunit(c *BaseunitContext)

	// EnterDerivedunit is called when entering the derivedunit production.
	EnterDerivedunit(c *DerivedunitContext)

	// ExitUom is called when exiting the uom production.
	ExitUom(c *UomContext)

	// ExitMeasure is called when exiting the measure production.
	ExitMeasure(c *MeasureContext)

	// ExitExponent is called when exiting the exponent production.
	ExitExponent(c *ExponentContext)

	// ExitPrefix is called when exiting the prefix production.
	ExitPrefix(c *PrefixContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitBaseunit is called when exiting the baseunit production.
	ExitBaseunit(c *BaseunitContext)

	// ExitDerivedunit is called when exiting the derivedunit production.
	ExitDerivedunit(c *DerivedunitContext)
}
