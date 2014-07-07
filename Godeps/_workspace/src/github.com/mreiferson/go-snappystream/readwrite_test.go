package snappystream

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestReaderWriter(t *testing.T) {
	var buf bytes.Buffer

	in := []byte("test")

	w := NewWriter(&buf)
	r := NewReader(&buf, VerifyChecksum)

	n, err := w.Write(in)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if n != len(in) {
		t.Fatalf("wrote wrong amount %d != %d", n, len(in))
	}

	out := make([]byte, len(in))
	n, err = r.Read(out)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if n != len(in) {
		t.Fatalf("read wrong amount %d != %d", n, len(in))
	}

	if !bytes.Equal(out, in) {
		t.Fatalf("bytes not equal %v != %v", out, in)
	}
}

func TestWriterChunk(t *testing.T) {
	var buf bytes.Buffer

	in := make([]byte, 128000)

	w := NewWriter(&buf)
	r := NewReader(&buf, VerifyChecksum)

	n, err := w.Write(in)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if n != len(in) {
		t.Fatalf("wrote wrong amount %d != %d", n, len(in))
	}

	out := make([]byte, len(in))
	n, err = io.ReadFull(r, out)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if n != len(in) {
		t.Fatalf("read wrong amount %d != %d", n, len(in))
	}

	if !bytes.Equal(out, in) {
		t.Fatalf("bytes not equal %v != %v", out, in)
	}
}

var testData = []byte(`
.TH XARGS 1L \" -*- nroff -*-
.SH NAME
xargs \- build and execute command lines from standard input
.SH SYNOPSIS
.B xargs
[\-0prtx] [\-e[eof-str]] [\-i[replace-str]] [\-l[max-lines]]
[\-n max-args] [\-s max-chars] [\-P max-procs] [\-\-null] [\-\-eof[=eof-str]]
[\-\-replace[=replace-str]] [\-\-max-lines[=max-lines]] [\-\-interactive]
[\-\-max-chars=max-chars] [\-\-verbose] [\-\-exit] [\-\-max-procs=max-procs]
[\-\-max-args=max-args] [\-\-no-run-if-empty] [\-\-version] [\-\-help]
[command [initial-arguments]]
.SH DESCRIPTION
This manual page
documents the GNU version of
.BR xargs .
.B xargs
reads arguments from the standard input, delimited by blanks (which can be
protected with double or single quotes or a backslash) or newlines,
and executes the
.I command
(default is /bin/echo) one or more times with any
.I initial-arguments
followed by arguments read from standard input.  Blank lines on the
standard input are ignored.
.P
.B xargs
exits with the following status:
.nf
0 if it succeeds
123 if any invocation of the command exited with status 1-125
124 if the command exited with status 255
125 if the command is killed by a signal
126 if the command cannot be run
127 if the command is not found
1 if some other error occurred.
.fi
.SS OPTIONS
.TP
.I "\-\-null, \-0"
Input filenames are terminated by a null character instead of by
whitespace, and the quotes and backslash are not special (every
character is taken literally).  Disables the end of file string, which
is treated like any other argument.  Useful when arguments might
contain white space, quote marks, or backslashes.  The GNU find
\-print0 option produces input suitable for this mode.
.TP
.I "\-\-eof[=eof-str], \-e[eof-str]"
Set the end of file string to \fIeof-str\fR.  If the end of file
string occurs as a line of input, the rest of the input is ignored.
If \fIeof-str\fR is omitted, there is no end of file string.  If this
option is not given, the end of file string defaults to "_".
.TP
.I "\-\-help"
Print a summary of the options to
.B xargs
and exit.
.TP
.I "\-\-replace[=replace-str], \-i[replace-str]"
Replace occurences of \fIreplace-str\fR in the initial arguments with
names read from standard input.
Also, unquoted blanks do not terminate arguments.
If \fIreplace-str\fR is omitted, it
defaults to "{}" (like for 'find \-exec').  Implies \fI\-x\fP and
\fI\-l 1\fP.
.TP
.I "\-\-max-lines[=max-lines], -l[max-lines]"
Use at most \fImax-lines\fR nonblank input lines per command line;
\fImax-lines\fR defaults to 1 if omitted.  Trailing blanks cause an
input line to be logically continued on the next input line.  Implies
\fI\-x\fR.
.TP
.I "\-\-max-args=max-args, \-n max-args"
Use at most \fImax-args\fR arguments per command line.  Fewer than
\fImax-args\fR arguments will be used if the size (see the \-s option)
is exceeded, unless the \-x option is given, in which case \fBxargs\fR
will exit.
.TP
.I "\-\-interactive, \-p"
Prompt the user about whether to run each command line and read a line
from the terminal.  Only run the command line if the response starts
with 'y' or 'Y'.  Implies \fI\-t\fR.
.TP
.I "\-\-no-run-if-empty, \-r"
If the standard input does not contain any nonblanks, do not run the
command.  Normally, the command is run once even if there is no input.
.TP
.I "\-\-max-chars=max-chars, \-s max-chars"
Use at most \fImax-chars\fR characters per command line, including the
command and initial arguments and the terminating nulls at the ends of
the argument strings.  The default is as large as possible, up to 20k
characters.
.TP
.I "\-\-verbose, \-t"
Print the command line on the standard error output before executing
it.
.TP
.I "\-\-version"
Print the version number of
.B xargs
and exit.
.TP
.I "\-\-exit, \-x"
Exit if the size (see the \fI\-s\fR option) is exceeded.
.TP
.I "\-\-max-procs=max-procs, \-P max-procs"
Run up to \fImax-procs\fR processes at a time; the default is 1.  If
\fImax-procs\fR is 0, \fBxargs\fR will run as many processes as
possible at a time.  Use the \fI\-n\fR option with \fI\-P\fR;
otherwise chances are that only one exec will be done.
.SH "SEE ALSO"
\fBfind\fP(1L), \fBlocate\fP(1L), \fBlocatedb\fP(5L), \fBupdatedb\fP(1)
\fBFinding Files\fP (on-line in Info, or printed)`)

func BenchmarkWriter(b *testing.B) {
	b.SetBytes(int64(len(testData)))
	w := NewWriter(ioutil.Discard)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n, err := w.Write(testData)
		if err != nil {
			b.Fatalf(err.Error())
		}
		if n != len(testData) {
			b.Fatalf("wrote wrong amount %d != %d", n, len(testData))
		}
	}
	b.StopTimer()
}
