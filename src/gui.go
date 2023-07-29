package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var (
	maxSumInput          *widgets.QDoubleSpinBox
	freeSumInput         *widgets.QDoubleSpinBox
	childSumInput        *widgets.QDoubleSpinBox
	werbeKostenLumpInput *widgets.QDoubleSpinBox
	socialMultiInput     *widgets.QDoubleSpinBox
	multiplierInput      *widgets.QDoubleSpinBox
	workIncomeInput      *widgets.QDoubleSpinBox
	familyIncomeInput    *widgets.QDoubleSpinBox
	socialIncomeInput    *widgets.QDoubleSpinBox
	otherIncomeInput     *widgets.QDoubleSpinBox
	assetsInput          *widgets.QDoubleSpinBox
	balanceInput         *widgets.QDoubleSpinBox
	childrenInput        *widgets.QSpinBox
	fundingRateField     *widgets.QLineEdit
)

func drawWindow() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Gorechner")

	// Layouts
	vParentLayout := widgets.NewQVBoxLayout() // übergeordnete Box
	hSubLayout := widgets.NewQHBoxLayout()
	vSubSubLayout := widgets.NewQVBoxLayout()
	zeroMargin := core.NewQMargins()
	vSubSubLayout.SetContentsMargins2(zeroMargin)
	vCalcLayout := widgets.NewQFormLayout(nil) // Calc-Teil
	vCalcLayout.SetRowWrapPolicy(2)
	vCalcLayout.SetVerticalSpacing(5)
	vIncomeLayout := widgets.NewQFormLayout(nil) // Income-Teil
	vIncomeLayout.SetRowWrapPolicy(2)
	vIncomeLayout.SetVerticalSpacing(5)
	vExtraLayout := widgets.NewQFormLayout(nil) // Extra-Teil
	vExtraLayout.SetRowWrapPolicy(2)
	vExtraLayout.SetVerticalSpacing(5)
	vFinalFundingLayout := widgets.NewQHBoxLayout()

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(vParentLayout)
	subWidget := widgets.NewQWidget(widget, 0)
	subWidget.SetLayout(hSubLayout)
	vParentLayout.AddWidget(subWidget, 0, 0)
	subSubWidget := widgets.NewQWidget(subWidget, 0)
	subSubWidget.SetLayout(vSubSubLayout)
	subSubWidget.SetContentsMargins2(zeroMargin)
	hSubLayout.AddWidget(subSubWidget, 0, 0)

	// Calc Gruppe
	calcGroup := widgets.NewQGroupBox2("BAföG-Festbeträge", nil)
	calcGroup.SetLayout(vCalcLayout)
	vSubSubLayout.AddWidget(calcGroup, 0, 0)

	// Felder für die CalcGruppe
	maxSumInput = newQDoubleSpinBoxWithMaximum(99999.)
	maxSumInput.SetValue(930.00)
	freeSumInput = newQDoubleSpinBoxWithMaximum(99999.)
	freeSumInput.SetValue(330.00)
	childSumInput = newQDoubleSpinBoxWithMaximum(99999.)
	childSumInput.SetValue(730.00)
	werbeKostenLumpInput = newQDoubleSpinBoxWithMaximum(99999.)
	werbeKostenLumpInput.SetValue(600.00)
	socialMultiInput = widgets.NewQDoubleSpinBox(nil)
	socialMultiInput.SetMaximum(100)
	socialMultiInput.SetDecimals(3)
	socialMultiInput.SetValue(21.6)
	socialMultiInput.SetSuffix("%")

	vCalcLayout.AddRow3("Höchstfördersatz", maxSumInput)
	vCalcLayout.AddRow3("Eigenfreibetrag", freeSumInput)
	vCalcLayout.AddRow3("Kinderfreibetrag", childSumInput)
	vCalcLayout.AddRow3("Werbekostenpauschale", werbeKostenLumpInput)
	vCalcLayout.AddRow3("Sozialpauschale", socialMultiInput)

	// Extra Gruppe
	extraGroup := widgets.NewQGroupBox2("Weitere Angaben", nil)
	extraGroup.SetLayout(vExtraLayout)
	vSubSubLayout.AddWidget(extraGroup, 0, 0)

	// Felder für die Extragruppe
	childrenInput = widgets.NewQSpinBox(nil)
	childrenInput.SetMaximum(30)
	childrenInput.SetValue(0)
	multiplierInput = widgets.NewQDoubleSpinBox(nil)
	multiplierInput.SetMaximum(1)
	multiplierInput.SetValue(1)
	multiplierInput.SetSingleStep(.01)
	multiplierInput.SetDecimals(15)
	vExtraLayout.AddRow3("Kinder", childrenInput)
	vExtraLayout.AddRow3("Multiplikator", multiplierInput)

	// Einkommensgruppe
	incomeGroup := widgets.NewQGroupBox2("Einkommensinformationen", nil)
	incomeGroup.SetLayout(vIncomeLayout)
	hSubLayout.AddWidget(incomeGroup, 0, 0)

	// Felder für die Einkommensgruppe
	workIncomeInput = newQDoubleSpinBoxWithMaximum(999999)
	familyIncomeInput = newQDoubleSpinBoxWithMaximum(999999)
	socialIncomeInput = newQDoubleSpinBoxWithMaximum(999999)
	otherIncomeInput = newQDoubleSpinBoxWithMaximum(999999)
	assetsInput = newQDoubleSpinBoxWithMaximum(999999)
	assetsInput.SetMinimum(-999999)
	balanceInput = newQDoubleSpinBoxWithMaximum(9999999)
	balanceInput.SetMinimum(-999999)
	vIncomeLayout.AddRow3("Einkommen aus Arbeit", workIncomeInput)
	vIncomeLayout.AddRow3("Einkommen durch Angehörige", familyIncomeInput)
	vIncomeLayout.AddRow3("Sozialleistungen etc.", socialIncomeInput)
	vIncomeLayout.AddRow3("Anderes Einkommen", otherIncomeInput)
	vIncomeLayout.AddRow3("Vermögen und Rücklagen", assetsInput)
	vIncomeLayout.AddRow3("Kontostand", balanceInput)

	// Förderraten Gruppe
	finalFundingGroup := widgets.NewQGroupBox(nil)
	finalFundingGroup.SetLayout(vFinalFundingLayout)
	vParentLayout.AddWidget(finalFundingGroup, 0, 0)

	// Labels und Felder für die Förderraten-Gruppe
	fundingRateLabel := widgets.NewQLabel2("<b>Mntl. Fördersumme:</b>", nil, 0)
	fundingRateLabel.SetTextFormat(2)
	vFinalFundingLayout.AddWidget(fundingRateLabel, 0, 0)
	fundingRateField = widgets.NewQLineEdit(nil)
	fundingRateField.SetReadOnly(true)
	fundingRateField.SetPlaceholderText("Fördersumme")
	fundingRateField.SetAlignment(0x0002 | 0x0080)
	fundingRateField.SetStyleSheet("font-weight: bold;")
	vFinalFundingLayout.AddWidget(fundingRateField, 0, 0)
	calculateButton := widgets.NewQPushButton2("Berechnen", nil)
	calculateButton.ConnectClicked(func(b bool) { updateFundingRateField() })
	vFinalFundingLayout.AddWidget(calculateButton, 0, 0)

	window.ConnectKeyPressEvent(func(event *gui.QKeyEvent) {
		switch key := event.Key(); key {
		case 0x01000000:
			app.Quit()
		case 0x01000005:
			updateFundingRateField()
		case 0x01000004:
			updateFundingRateField()
		}
	})

	window.SetCentralWidget(widget)
	window.Show()
	widgets.QApplication_Exec()
}

func newQDoubleSpinBoxWithMaximum(max float64) *widgets.QDoubleSpinBox {
	spinBox := widgets.NewQDoubleSpinBox(nil)
	spinBox.SetMaximum(max)
	spinBox.SetSuffix("€")
	return spinBox
}

func updateFundingRateField() {
	income := incomeInfo{
		workIncome:   workIncomeInput.Value(),
		familyIncome: familyIncomeInput.Value(),
		socialIncome: socialIncomeInput.Value(),
		otherIncome:  otherIncomeInput.Value(),
		assets:       assetsInput.Value(),
		balance:      balanceInput.Value(),
		children:     childrenInput.Value(),
	}

	calc := calcInfo{
		maxSum:          maxSumInput.Value(),
		freeSum:         freeSumInput.Value(),
		childSum:        childSumInput.Value(),
		werbeKostenLump: werbeKostenLumpInput.Value(),
		socialMulti:     socialMultiInput.Value() / 100,
		multiplier:      multiplierInput.Value(),
	}

	qLocale := fundingRateField.Locale()
	text := qLocale.ToCurrencyString7(getFundingRate(income, calc), qLocale.CurrencySymbol(1))
	fundingRateField.SetText(text)
}
