#include "textflag.h"

#define round(x, y, z, t, u, v) \
	MOVO	x, t		\
	PSLLL	$24, x		\
	PSRLL	$8, t		\
	PXOR	t, x		\
				\
	MOVO	y, u		\
	PSLLL	$9, y		\
	PSRLL	$23, u		\
	PXOR	u, y		\
				\
	MOVO	x, t		\
	PAND	y, t		\
	PSLLL	$3, t		\
	PXOR	y, t		\
	PXOR	z, t		\
				\
	MOVO	x, u		\
	POR	z, u		\
	PSLLL	$1, u		\
	PXOR	x, u		\
	PXOR	y, u		\
				\
	MOVO	y, v		\
	PAND	z, v		\
	PSLLL	$2, v		\
	PSLLL	$1, z		\
	PXOR	z, v		\
	PXOR	x, v		\

DATA	round_constant<>+0(SB)/8, $0x9e377918
DATA	round_constant<>+8(SB)/8, $0
GLOBL	round_constant<>(SB), (NOPTR+RODATA), $16

DATA	counter<>+0(SB)/8, $4
DATA	counter<>+8(SB)/8, $0
GLOBL	counter<>(SB), (NOPTR+RODATA), $16

TEXT	·permuteSSE2(SB), $0-4
	MOVL	s+0(FP), AX

	MOVOU	(AX), X0
	MOVOU	16(AX), X1
	MOVOU	32(AX), X2

	MOVL	$24, CX

	MOVO	round_constant<>(SB), X6
	MOVO	counter<>(SB), X7

loop:
	SUBL	$4, CX

	round(X0, X1, X2, X3, X4, X5)
	PSHUFL	$0xB1, X3, X3	// small swap: 10 11 00 01

	// round constant
	PXOR	X6, X3
	PSUBL	X7, X6

	round(X3, X4, X5, X0, X1, X2)
	round(X0, X1, X2, X3, X4, X5)

	PSHUFL $0x4E, X3, X3	// big swap: 01 00 11 10

	round(X3, X4, X5, X0, X1, X2)

	TESTL	CX, CX
	JGT	$0, loop

	MOVOU	X0, (AX)
	MOVOU	X1, 16(AX)
	MOVOU	X2, 32(AX)

	RET

TEXT	·supportSSE2(SB), NOSPLIT, $0-1
	MOVL	$1, AX
	CPUID
	SHRL	$26, DX
	ANDL	$1, DX        // DX != 0 if support SSE2
	MOVB	DX, ret+0(FP)
	RET
