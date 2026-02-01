package ui

import (
	"fmt"
	"io"
	"time"

	"github.com/briandowns/spinner"
)

const (
	spinnerSpeed    = 80 * time.Millisecond
	charset         = 14                     // Arc - animação suave
	stepMinDuration = 600 * time.Millisecond // Tempo mínimo para ver o spinner
	stepPause       = 180 * time.Millisecond // Pausa entre steps
)

type Stepper struct {
	out     io.Writer
	current int
	total   int
	s       *spinner.Spinner
}

func NewStepper(out io.Writer, total int) *Stepper {
	s := spinner.New(spinner.CharSets[charset], spinnerSpeed)
	s.Writer = out
	s.Color("cyan")
	return &Stepper{
		out:   out,
		total: total,
		s:     s,
	}
}

func (st *Stepper) StartStep(msg string) {
	time.Sleep(stepPause)
	st.current++
	st.s.Prefix = fmt.Sprintf("  [%d/%d] ", st.current, st.total)
	st.s.Suffix = " " + msg
	st.s.Start()
}

func (st *Stepper) UpdateAndContinue(msg string) {
	st.s.Suffix = " " + msg
}

func (st *Stepper) Ok(msg string, elapsed time.Duration) {
	st.s.Stop()
	remaining := stepMinDuration - elapsed
	if remaining > 0 {
		time.Sleep(remaining)
	}
	fmt.Fprintf(st.out, "  %s %s %s\n",
		styleStepNum.Render(fmt.Sprintf("[%d/%d]", st.current, st.total)),
		styleOk.Render("✔"),
		styleMessage.Render(msg),
	)
}

func (st *Stepper) Skip(msg string) {
	st.s.Stop()
	time.Sleep(stepPause)
	fmt.Fprintf(st.out, "  %s %s %s\n",
		styleStepNum.Render(fmt.Sprintf("[%d/%d]", st.current, st.total)),
		styleSkip.Render("○"),
		styleMessage.Render(msg),
	)
}

func (st *Stepper) Warn(msg string) {
	st.s.Stop()
	fmt.Fprintf(st.out, "  %s %s %s\n",
		styleStepNum.Render(fmt.Sprintf("[%d/%d]", st.current, st.total)),
		styleWarn.Render("⚠"),
		styleMessage.Render(msg),
	)
}

func (st *Stepper) Fail(msg string) {
	st.s.Stop()
	fmt.Fprintf(st.out, "  %s %s %s\n",
		styleStepNum.Render(fmt.Sprintf("[%d/%d]", st.current, st.total)),
		styleFail.Render("✗"),
		styleMessage.Render(msg),
	)
}
