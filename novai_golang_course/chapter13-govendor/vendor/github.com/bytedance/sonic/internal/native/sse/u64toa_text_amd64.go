// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package sse

var _text_u64toa = []byte{
	// .p2align 4, 0x00
	// LCPI0_0
	0x59, 0x17, 0xb7, 0xd1, 0x00, 0x00, 0x00, 0x00, // .quad 3518437209
	0x59, 0x17, 0xb7, 0xd1, 0x00, 0x00, 0x00, 0x00, //0x00000008 .quad 3518437209
	//0x00000010 LCPI0_1
	0xc5, 0x20, //0x00000010 .word 8389
	0x7b, 0x14, //0x00000012 .word 5243
	0x34, 0x33, //0x00000014 .word 13108
	0x00, 0x80, //0x00000016 .word 32768
	0xc5, 0x20, //0x00000018 .word 8389
	0x7b, 0x14, //0x0000001a .word 5243
	0x34, 0x33, //0x0000001c .word 13108
	0x00, 0x80, //0x0000001e .word 32768
	//0x00000020 LCPI0_2
	0x80, 0x00, //0x00000020 .word 128
	0x00, 0x08, //0x00000022 .word 2048
	0x00, 0x20, //0x00000024 .word 8192
	0x00, 0x80, //0x00000026 .word 32768
	0x80, 0x00, //0x00000028 .word 128
	0x00, 0x08, //0x0000002a .word 2048
	0x00, 0x20, //0x0000002c .word 8192
	0x00, 0x80, //0x0000002e .word 32768
	//0x00000030 LCPI0_3
	0x0a, 0x00, //0x00000030 .word 10
	0x0a, 0x00, //0x00000032 .word 10
	0x0a, 0x00, //0x00000034 .word 10
	0x0a, 0x00, //0x00000036 .word 10
	0x0a, 0x00, //0x00000038 .word 10
	0x0a, 0x00, //0x0000003a .word 10
	0x0a, 0x00, //0x0000003c .word 10
	0x0a, 0x00, //0x0000003e .word 10
	//0x00000040 LCPI0_4
	0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, //0x00000040 QUAD $0x3030303030303030; QUAD $0x3030303030303030  // .space 16, '0000000000000000'
	//0x00000050 .p2align 4, 0x90
	//0x00000050 _u64toa
	0x55, //0x00000050 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000051 movq         %rsp, %rbp
	0x48, 0x81, 0xfe, 0x0f, 0x27, 0x00, 0x00, //0x00000054 cmpq         $9999, %rsi
	0x0f, 0x87, 0xa2, 0x00, 0x00, 0x00, //0x0000005b ja           LBB0_8
	0x0f, 0xb7, 0xc6, //0x00000061 movzwl       %si, %eax
	0xc1, 0xe8, 0x02, //0x00000064 shrl         $2, %eax
	0x69, 0xc0, 0x7b, 0x14, 0x00, 0x00, //0x00000067 imull        $5243, %eax, %eax
	0xc1, 0xe8, 0x11, //0x0000006d shrl         $17, %eax
	0x48, 0x8d, 0x14, 0x00, //0x00000070 leaq         (%rax,%rax), %rdx
	0x6b, 0xc0, 0x64, //0x00000074 imull        $100, %eax, %eax
	0x89, 0xf1, //0x00000077 movl         %esi, %ecx
	0x29, 0xc1, //0x00000079 subl         %eax, %ecx
	0x0f, 0xb7, 0xc1, //0x0000007b movzwl       %cx, %eax
	0x48, 0x01, 0xc0, //0x0000007e addq         %rax, %rax
	0x81, 0xfe, 0xe8, 0x03, 0x00, 0x00, //0x00000081 cmpl         $1000, %esi
	0x0f, 0x82, 0x16, 0x00, 0x00, 0x00, //0x00000087 jb           LBB0_3
	0x48, 0x8d, 0x0d, 0xac, 0x04, 0x00, 0x00, //0x0000008d leaq         $1196(%rip), %rcx  /* _Digits+0(%rip) */
	0x8a, 0x0c, 0x0a, //0x00000094 movb         (%rdx,%rcx), %cl
	0x88, 0x0f, //0x00000097 movb         %cl, (%rdi)
	0xb9, 0x01, 0x00, 0x00, 0x00, //0x00000099 movl         $1, %ecx
	0xe9, 0x0b, 0x00, 0x00, 0x00, //0x0000009e jmp          LBB0_4
	//0x000000a3 LBB0_3
	0x31, 0xc9, //0x000000a3 xorl         %ecx, %ecx
	0x83, 0xfe, 0x64, //0x000000a5 cmpl         $100, %esi
	0x0f, 0x82, 0x45, 0x00, 0x00, 0x00, //0x000000a8 jb           LBB0_5
	//0x000000ae LBB0_4
	0x0f, 0xb7, 0xd2, //0x000000ae movzwl       %dx, %edx
	0x48, 0x83, 0xca, 0x01, //0x000000b1 orq          $1, %rdx
	0x48, 0x8d, 0x35, 0x84, 0x04, 0x00, 0x00, //0x000000b5 leaq         $1156(%rip), %rsi  /* _Digits+0(%rip) */
	0x8a, 0x14, 0x32, //0x000000bc movb         (%rdx,%rsi), %dl
	0x89, 0xce, //0x000000bf movl         %ecx, %esi
	0xff, 0xc1, //0x000000c1 incl         %ecx
	0x88, 0x14, 0x37, //0x000000c3 movb         %dl, (%rdi,%rsi)
	//0x000000c6 LBB0_6
	0x48, 0x8d, 0x15, 0x73, 0x04, 0x00, 0x00, //0x000000c6 leaq         $1139(%rip), %rdx  /* _Digits+0(%rip) */
	0x8a, 0x14, 0x10, //0x000000cd movb         (%rax,%rdx), %dl
	0x89, 0xce, //0x000000d0 movl         %ecx, %esi
	0xff, 0xc1, //0x000000d2 incl         %ecx
	0x88, 0x14, 0x37, //0x000000d4 movb         %dl, (%rdi,%rsi)
	//0x000000d7 LBB0_7
	0x0f, 0xb7, 0xc0, //0x000000d7 movzwl       %ax, %eax
	0x48, 0x83, 0xc8, 0x01, //0x000000da orq          $1, %rax
	0x48, 0x8d, 0x15, 0x5b, 0x04, 0x00, 0x00, //0x000000de leaq         $1115(%rip), %rdx  /* _Digits+0(%rip) */
	0x8a, 0x04, 0x10, //0x000000e5 movb         (%rax,%rdx), %al
	0x89, 0xca, //0x000000e8 movl         %ecx, %edx
	0xff, 0xc1, //0x000000ea incl         %ecx
	0x88, 0x04, 0x17, //0x000000ec movb         %al, (%rdi,%rdx)
	0x89, 0xc8, //0x000000ef movl         %ecx, %eax
	0x5d, //0x000000f1 popq         %rbp
	0xc3, //0x000000f2 retq         
	//0x000000f3 LBB0_5
	0x31, 0xc9, //0x000000f3 xorl         %ecx, %ecx
	0x83, 0xfe, 0x0a, //0x000000f5 cmpl         $10, %esi
	0x0f, 0x83, 0xc8, 0xff, 0xff, 0xff, //0x000000f8 jae          LBB0_6
	0xe9, 0xd4, 0xff, 0xff, 0xff, //0x000000fe jmp          LBB0_7
	//0x00000103 LBB0_8
	0x48, 0x81, 0xfe, 0xff, 0xe0, 0xf5, 0x05, //0x00000103 cmpq         $99999999, %rsi
	0x0f, 0x87, 0x1e, 0x01, 0x00, 0x00, //0x0000010a ja           LBB0_16
	0x89, 0xf0, //0x00000110 movl         %esi, %eax
	0xba, 0x59, 0x17, 0xb7, 0xd1, //0x00000112 movl         $3518437209, %edx
	0x48, 0x0f, 0xaf, 0xd0, //0x00000117 imulq        %rax, %rdx
	0x48, 0xc1, 0xea, 0x2d, //0x0000011b shrq         $45, %rdx
	0x44, 0x69, 0xc2, 0x10, 0x27, 0x00, 0x00, //0x0000011f imull        $10000, %edx, %r8d
	0x89, 0xf1, //0x00000126 movl         %esi, %ecx
	0x44, 0x29, 0xc1, //0x00000128 subl         %r8d, %ecx
	0x4c, 0x69, 0xd0, 0x83, 0xde, 0x1b, 0x43, //0x0000012b imulq        $1125899907, %rax, %r10
	0x49, 0xc1, 0xea, 0x31, //0x00000132 shrq         $49, %r10
	0x41, 0x83, 0xe2, 0xfe, //0x00000136 andl         $-2, %r10d
	0x0f, 0xb7, 0xc2, //0x0000013a movzwl       %dx, %eax
	0xc1, 0xe8, 0x02, //0x0000013d shrl         $2, %eax
	0x69, 0xc0, 0x7b, 0x14, 0x00, 0x00, //0x00000140 imull        $5243, %eax, %eax
	0xc1, 0xe8, 0x11, //0x00000146 shrl         $17, %eax
	0x6b, 0xc0, 0x64, //0x00000149 imull        $100, %eax, %eax
	0x29, 0xc2, //0x0000014c subl         %eax, %edx
	0x44, 0x0f, 0xb7, 0xca, //0x0000014e movzwl       %dx, %r9d
	0x4d, 0x01, 0xc9, //0x00000152 addq         %r9, %r9
	0x0f, 0xb7, 0xc1, //0x00000155 movzwl       %cx, %eax
	0xc1, 0xe8, 0x02, //0x00000158 shrl         $2, %eax
	0x69, 0xc0, 0x7b, 0x14, 0x00, 0x00, //0x0000015b imull        $5243, %eax, %eax
	0xc1, 0xe8, 0x11, //0x00000161 shrl         $17, %eax
	0x4c, 0x8d, 0x04, 0x00, //0x00000164 leaq         (%rax,%rax), %r8
	0x6b, 0xc0, 0x64, //0x00000168 imull        $100, %eax, %eax
	0x29, 0xc1, //0x0000016b subl         %eax, %ecx
	0x44, 0x0f, 0xb7, 0xd9, //0x0000016d movzwl       %cx, %r11d
	0x4d, 0x01, 0xdb, //0x00000171 addq         %r11, %r11
	0x81, 0xfe, 0x80, 0x96, 0x98, 0x00, //0x00000174 cmpl         $10000000, %esi
	0x0f, 0x82, 0x17, 0x00, 0x00, 0x00, //0x0000017a jb           LBB0_11
	0x48, 0x8d, 0x05, 0xb9, 0x03, 0x00, 0x00, //0x00000180 leaq         $953(%rip), %rax  /* _Digits+0(%rip) */
	0x41, 0x8a, 0x04, 0x02, //0x00000187 movb         (%r10,%rax), %al
	0x88, 0x07, //0x0000018b movb         %al, (%rdi)
	0xb9, 0x01, 0x00, 0x00, 0x00, //0x0000018d movl         $1, %ecx
	0xe9, 0x0e, 0x00, 0x00, 0x00, //0x00000192 jmp          LBB0_12
	//0x00000197 LBB0_11
	0x31, 0xc9, //0x00000197 xorl         %ecx, %ecx
	0x81, 0xfe, 0x40, 0x42, 0x0f, 0x00, //0x00000199 cmpl         $1000000, %esi
	0x0f, 0x82, 0x76, 0x00, 0x00, 0x00, //0x0000019f jb           LBB0_13
	//0x000001a5 LBB0_12
	0x44, 0x89, 0xd0, //0x000001a5 movl         %r10d, %eax
	0x48, 0x83, 0xc8, 0x01, //0x000001a8 orq          $1, %rax
	0x48, 0x8d, 0x35, 0x8d, 0x03, 0x00, 0x00, //0x000001ac leaq         $909(%rip), %rsi  /* _Digits+0(%rip) */
	0x8a, 0x04, 0x30, //0x000001b3 movb         (%rax,%rsi), %al
	0x89, 0xce, //0x000001b6 movl         %ecx, %esi
	0xff, 0xc1, //0x000001b8 incl         %ecx
	0x88, 0x04, 0x37, //0x000001ba movb         %al, (%rdi,%rsi)
	//0x000001bd LBB0_14
	0x48, 0x8d, 0x05, 0x7c, 0x03, 0x00, 0x00, //0x000001bd leaq         $892(%rip), %rax  /* _Digits+0(%rip) */
	0x41, 0x8a, 0x04, 0x01, //0x000001c4 movb         (%r9,%rax), %al
	0x89, 0xce, //0x000001c8 movl         %ecx, %esi
	0xff, 0xc1, //0x000001ca incl         %ecx
	0x88, 0x04, 0x37, //0x000001cc movb         %al, (%rdi,%rsi)
	//0x000001cf LBB0_15
	0x41, 0x0f, 0xb7, 0xc1, //0x000001cf movzwl       %r9w, %eax
	0x48, 0x83, 0xc8, 0x01, //0x000001d3 orq          $1, %rax
	0x48, 0x8d, 0x35, 0x62, 0x03, 0x00, 0x00, //0x000001d7 leaq         $866(%rip), %rsi  /* _Digits+0(%rip) */
	0x8a, 0x04, 0x30, //0x000001de movb         (%rax,%rsi), %al
	0x89, 0xca, //0x000001e1 movl         %ecx, %edx
	0x88, 0x04, 0x3a, //0x000001e3 movb         %al, (%rdx,%rdi)
	0x41, 0x8a, 0x04, 0x30, //0x000001e6 movb         (%r8,%rsi), %al
	0x88, 0x44, 0x3a, 0x01, //0x000001ea movb         %al, $1(%rdx,%rdi)
	0x41, 0x0f, 0xb7, 0xc0, //0x000001ee movzwl       %r8w, %eax
	0x48, 0x83, 0xc8, 0x01, //0x000001f2 orq          $1, %rax
	0x8a, 0x04, 0x30, //0x000001f6 movb         (%rax,%rsi), %al
	0x88, 0x44, 0x3a, 0x02, //0x000001f9 movb         %al, $2(%rdx,%rdi)
	0x41, 0x8a, 0x04, 0x33, //0x000001fd movb         (%r11,%rsi), %al
	0x88, 0x44, 0x3a, 0x03, //0x00000201 movb         %al, $3(%rdx,%rdi)
	0x41, 0x0f, 0xb7, 0xc3, //0x00000205 movzwl       %r11w, %eax
	0x48, 0x83, 0xc8, 0x01, //0x00000209 orq          $1, %rax
	0x8a, 0x04, 0x30, //0x0000020d movb         (%rax,%rsi), %al
	0x83, 0xc1, 0x05, //0x00000210 addl         $5, %ecx
	0x88, 0x44, 0x3a, 0x04, //0x00000213 movb         %al, $4(%rdx,%rdi)
	0x89, 0xc8, //0x00000217 movl         %ecx, %eax
	0x5d, //0x00000219 popq         %rbp
	0xc3, //0x0000021a retq         
	//0x0000021b LBB0_13
	0x31, 0xc9, //0x0000021b xorl         %ecx, %ecx
	0x81, 0xfe, 0xa0, 0x86, 0x01, 0x00, //0x0000021d cmpl         $100000, %esi
	0x0f, 0x83, 0x94, 0xff, 0xff, 0xff, //0x00000223 jae          LBB0_14
	0xe9, 0xa1, 0xff, 0xff, 0xff, //0x00000229 jmp          LBB0_15
	//0x0000022e LBB0_16
	0x48, 0xb8, 0xff, 0xff, 0xc0, 0x6f, 0xf2, 0x86, 0x23, 0x00, //0x0000022e movabsq      $9999999999999999, %rax
	0x48, 0x39, 0xc6, //0x00000238 cmpq         %rax, %rsi
	0x0f, 0x87, 0x15, 0x01, 0x00, 0x00, //0x0000023b ja           LBB0_18
	0x48, 0xb9, 0xfd, 0xce, 0x61, 0x84, 0x11, 0x77, 0xcc, 0xab, //0x00000241 movabsq      $-6067343680855748867, %rcx
	0x48, 0x89, 0xf0, //0x0000024b movq         %rsi, %rax
	0x48, 0xf7, 0xe1, //0x0000024e mulq         %rcx
	0x48, 0xc1, 0xea, 0x1a, //0x00000251 shrq         $26, %rdx
	0x69, 0xc2, 0x00, 0xe1, 0xf5, 0x05, //0x00000255 imull        $100000000, %edx, %eax
	0x29, 0xc6, //0x0000025b subl         %eax, %esi
	0x66, 0x0f, 0x6e, 0xc2, //0x0000025d movd         %edx, %xmm0
	0xf3, 0x0f, 0x6f, 0x0d, 0x97, 0xfd, 0xff, 0xff, //0x00000261 movdqu       $-617(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	0x66, 0x0f, 0x6f, 0xd0, //0x00000269 movdqa       %xmm0, %xmm2
	0x66, 0x0f, 0xf4, 0xd1, //0x0000026d pmuludq      %xmm1, %xmm2
	0x66, 0x0f, 0x73, 0xd2, 0x2d, //0x00000271 psrlq        $45, %xmm2
	0xb8, 0x10, 0x27, 0x00, 0x00, //0x00000276 movl         $10000, %eax
	0x66, 0x48, 0x0f, 0x6e, 0xd8, //0x0000027b movq         %rax, %xmm3
	0x66, 0x0f, 0x6f, 0xe2, //0x00000280 movdqa       %xmm2, %xmm4
	0x66, 0x0f, 0xf4, 0xe3, //0x00000284 pmuludq      %xmm3, %xmm4
	0x66, 0x0f, 0xfa, 0xc4, //0x00000288 psubd        %xmm4, %xmm0
	0x66, 0x0f, 0x61, 0xd0, //0x0000028c punpcklwd    %xmm0, %xmm2
	0x66, 0x0f, 0x73, 0xf2, 0x02, //0x00000290 psllq        $2, %xmm2
	0xf2, 0x0f, 0x70, 0xc2, 0x50, //0x00000295 pshuflw      $80, %xmm2, %xmm0
	0x66, 0x0f, 0x70, 0xc0, 0x50, //0x0000029a pshufd       $80, %xmm0, %xmm0
	0xf3, 0x0f, 0x6f, 0x15, 0x69, 0xfd, 0xff, 0xff, //0x0000029f movdqu       $-663(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	0x66, 0x0f, 0xe4, 0xc2, //0x000002a7 pmulhuw      %xmm2, %xmm0
	0xf3, 0x0f, 0x6f, 0x25, 0x6d, 0xfd, 0xff, 0xff, //0x000002ab movdqu       $-659(%rip), %xmm4  /* LCPI0_2+0(%rip) */
	0x66, 0x0f, 0xe4, 0xc4, //0x000002b3 pmulhuw      %xmm4, %xmm0
	0xf3, 0x0f, 0x6f, 0x2d, 0x71, 0xfd, 0xff, 0xff, //0x000002b7 movdqu       $-655(%rip), %xmm5  /* LCPI0_3+0(%rip) */
	0x66, 0x0f, 0x6f, 0xf0, //0x000002bf movdqa       %xmm0, %xmm6
	0x66, 0x0f, 0xd5, 0xf5, //0x000002c3 pmullw       %xmm5, %xmm6
	0x66, 0x0f, 0x73, 0xf6, 0x10, //0x000002c7 psllq        $16, %xmm6
	0x66, 0x0f, 0xf9, 0xc6, //0x000002cc psubw        %xmm6, %xmm0
	0x66, 0x0f, 0x6e, 0xf6, //0x000002d0 movd         %esi, %xmm6
	0x66, 0x0f, 0xf4, 0xce, //0x000002d4 pmuludq      %xmm6, %xmm1
	0x66, 0x0f, 0x73, 0xd1, 0x2d, //0x000002d8 psrlq        $45, %xmm1
	0x66, 0x0f, 0xf4, 0xd9, //0x000002dd pmuludq      %xmm1, %xmm3
	0x66, 0x0f, 0xfa, 0xf3, //0x000002e1 psubd        %xmm3, %xmm6
	0x66, 0x0f, 0x61, 0xce, //0x000002e5 punpcklwd    %xmm6, %xmm1
	0x66, 0x0f, 0x73, 0xf1, 0x02, //0x000002e9 psllq        $2, %xmm1
	0xf2, 0x0f, 0x70, 0xc9, 0x50, //0x000002ee pshuflw      $80, %xmm1, %xmm1
	0x66, 0x0f, 0x70, 0xc9, 0x50, //0x000002f3 pshufd       $80, %xmm1, %xmm1
	0x66, 0x0f, 0xe4, 0xca, //0x000002f8 pmulhuw      %xmm2, %xmm1
	0x66, 0x0f, 0xe4, 0xcc, //0x000002fc pmulhuw      %xmm4, %xmm1
	0x66, 0x0f, 0xd5, 0xe9, //0x00000300 pmullw       %xmm1, %xmm5
	0x66, 0x0f, 0x73, 0xf5, 0x10, //0x00000304 psllq        $16, %xmm5
	0x66, 0x0f, 0xf9, 0xcd, //0x00000309 psubw        %xmm5, %xmm1
	0x66, 0x0f, 0x67, 0xc1, //0x0000030d packuswb     %xmm1, %xmm0
	0xf3, 0x0f, 0x6f, 0x0d, 0x27, 0xfd, 0xff, 0xff, //0x00000311 movdqu       $-729(%rip), %xmm1  /* LCPI0_4+0(%rip) */
	0x66, 0x0f, 0xfc, 0xc8, //0x00000319 paddb        %xmm0, %xmm1
	0x66, 0x0f, 0xef, 0xd2, //0x0000031d pxor         %xmm2, %xmm2
	0x66, 0x0f, 0x74, 0xd0, //0x00000321 pcmpeqb      %xmm0, %xmm2
	0x66, 0x0f, 0xd7, 0xc2, //0x00000325 pmovmskb     %xmm2, %eax
	0x0d, 0x00, 0x80, 0x00, 0x00, //0x00000329 orl          $32768, %eax
	0x35, 0xff, 0x7f, 0xff, 0xff, //0x0000032e xorl         $-32769, %eax
	0x0f, 0xbc, 0xc0, //0x00000333 bsfl         %eax, %eax
	0xb9, 0x10, 0x00, 0x00, 0x00, //0x00000336 movl         $16, %ecx
	0x29, 0xc1, //0x0000033b subl         %eax, %ecx
	0x48, 0xc1, 0xe0, 0x04, //0x0000033d shlq         $4, %rax
	0x48, 0x8d, 0x15, 0xc8, 0x02, 0x00, 0x00, //0x00000341 leaq         $712(%rip), %rdx  /* _VecShiftShuffles+0(%rip) */
	0x66, 0x0f, 0x38, 0x00, 0x0c, 0x10, //0x00000348 pshufb       (%rax,%rdx), %xmm1
	0xf3, 0x0f, 0x7f, 0x0f, //0x0000034e movdqu       %xmm1, (%rdi)
	0x89, 0xc8, //0x00000352 movl         %ecx, %eax
	0x5d, //0x00000354 popq         %rbp
	0xc3, //0x00000355 retq         
	//0x00000356 LBB0_18
	0x48, 0xb9, 0x57, 0x78, 0x13, 0xb1, 0x2f, 0x65, 0xa5, 0x39, //0x00000356 movabsq      $4153837486827862103, %rcx
	0x48, 0x89, 0xf0, //0x00000360 movq         %rsi, %rax
	0x48, 0xf7, 0xe1, //0x00000363 mulq         %rcx
	0x48, 0xc1, 0xea, 0x33, //0x00000366 shrq         $51, %rdx
	0x48, 0xb8, 0x00, 0x00, 0xc1, 0x6f, 0xf2, 0x86, 0x23, 0x00, //0x0000036a movabsq      $10000000000000000, %rax
	0x48, 0x0f, 0xaf, 0xc2, //0x00000374 imulq        %rdx, %rax
	0x48, 0x29, 0xc6, //0x00000378 subq         %rax, %rsi
	0x83, 0xfa, 0x09, //0x0000037b cmpl         $9, %edx
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x0000037e ja           LBB0_20
	0x80, 0xc2, 0x30, //0x00000384 addb         $48, %dl
	0x88, 0x17, //0x00000387 movb         %dl, (%rdi)
	0xb9, 0x01, 0x00, 0x00, 0x00, //0x00000389 movl         $1, %ecx
	0xe9, 0xba, 0x00, 0x00, 0x00, //0x0000038e jmp          LBB0_25
	//0x00000393 LBB0_20
	0x83, 0xfa, 0x63, //0x00000393 cmpl         $99, %edx
	0x0f, 0x87, 0x1f, 0x00, 0x00, 0x00, //0x00000396 ja           LBB0_22
	0x89, 0xd0, //0x0000039c movl         %edx, %eax
	0x48, 0x8d, 0x0d, 0x9b, 0x01, 0x00, 0x00, //0x0000039e leaq         $411(%rip), %rcx  /* _Digits+0(%rip) */
	0x8a, 0x14, 0x41, //0x000003a5 movb         (%rcx,%rax,2), %dl
	0x8a, 0x44, 0x41, 0x01, //0x000003a8 movb         $1(%rcx,%rax,2), %al
	0x88, 0x17, //0x000003ac movb         %dl, (%rdi)
	0x88, 0x47, 0x01, //0x000003ae movb         %al, $1(%rdi)
	0xb9, 0x02, 0x00, 0x00, 0x00, //0x000003b1 movl         $2, %ecx
	0xe9, 0x92, 0x00, 0x00, 0x00, //0x000003b6 jmp          LBB0_25
	//0x000003bb LBB0_22
	0x89, 0xd0, //0x000003bb movl         %edx, %eax
	0xc1, 0xe8, 0x02, //0x000003bd shrl         $2, %eax
	0x69, 0xc0, 0x7b, 0x14, 0x00, 0x00, //0x000003c0 imull        $5243, %eax, %eax
	0xc1, 0xe8, 0x11, //0x000003c6 shrl         $17, %eax
	0x81, 0xfa, 0xe7, 0x03, 0x00, 0x00, //0x000003c9 cmpl         $999, %edx
	0x0f, 0x87, 0x3c, 0x00, 0x00, 0x00, //0x000003cf ja           LBB0_24
	0x83, 0xc0, 0x30, //0x000003d5 addl         $48, %eax
	0x88, 0x07, //0x000003d8 movb         %al, (%rdi)
	0x0f, 0xb7, 0xc2, //0x000003da movzwl       %dx, %eax
	0x89, 0xc1, //0x000003dd movl         %eax, %ecx
	0xc1, 0xe9, 0x02, //0x000003df shrl         $2, %ecx
	0x69, 0xc9, 0x7b, 0x14, 0x00, 0x00, //0x000003e2 imull        $5243, %ecx, %ecx
	0xc1, 0xe9, 0x11, //0x000003e8 shrl         $17, %ecx
	0x6b, 0xc9, 0x64, //0x000003eb imull        $100, %ecx, %ecx
	0x29, 0xc8, //0x000003ee subl         %ecx, %eax
	0x0f, 0xb7, 0xc0, //0x000003f0 movzwl       %ax, %eax
	0x48, 0x8d, 0x0d, 0x46, 0x01, 0x00, 0x00, //0x000003f3 leaq         $326(%rip), %rcx  /* _Digits+0(%rip) */
	0x8a, 0x14, 0x41, //0x000003fa movb         (%rcx,%rax,2), %dl
	0x8a, 0x44, 0x41, 0x01, //0x000003fd movb         $1(%rcx,%rax,2), %al
	0x88, 0x57, 0x01, //0x00000401 movb         %dl, $1(%rdi)
	0x88, 0x47, 0x02, //0x00000404 movb         %al, $2(%rdi)
	0xb9, 0x03, 0x00, 0x00, 0x00, //0x00000407 movl         $3, %ecx
	0xe9, 0x3c, 0x00, 0x00, 0x00, //0x0000040c jmp          LBB0_25
	//0x00000411 LBB0_24
	0x6b, 0xc8, 0x64, //0x00000411 imull        $100, %eax, %ecx
	0x29, 0xca, //0x00000414 subl         %ecx, %edx
	0x0f, 0xb7, 0xc0, //0x00000416 movzwl       %ax, %eax
	0x4c, 0x8d, 0x05, 0x20, 0x01, 0x00, 0x00, //0x00000419 leaq         $288(%rip), %r8  /* _Digits+0(%rip) */
	0x41, 0x8a, 0x0c, 0x40, //0x00000420 movb         (%r8,%rax,2), %cl
	0x41, 0x8a, 0x44, 0x40, 0x01, //0x00000424 movb         $1(%r8,%rax,2), %al
	0x88, 0x0f, //0x00000429 movb         %cl, (%rdi)
	0x88, 0x47, 0x01, //0x0000042b movb         %al, $1(%rdi)
	0x0f, 0xb7, 0xc2, //0x0000042e movzwl       %dx, %eax
	0x41, 0x8a, 0x0c, 0x40, //0x00000431 movb         (%r8,%rax,2), %cl
	0x48, 0x01, 0xc0, //0x00000435 addq         %rax, %rax
	0x88, 0x4f, 0x02, //0x00000438 movb         %cl, $2(%rdi)
	0x83, 0xc8, 0x01, //0x0000043b orl          $1, %eax
	0x0f, 0xb7, 0xc0, //0x0000043e movzwl       %ax, %eax
	0x42, 0x8a, 0x04, 0x00, //0x00000441 movb         (%rax,%r8), %al
	0x88, 0x47, 0x03, //0x00000445 movb         %al, $3(%rdi)
	0xb9, 0x04, 0x00, 0x00, 0x00, //0x00000448 movl         $4, %ecx
	//0x0000044d LBB0_25
	0x48, 0xba, 0xfd, 0xce, 0x61, 0x84, 0x11, 0x77, 0xcc, 0xab, //0x0000044d movabsq      $-6067343680855748867, %rdx
	0x48, 0x89, 0xf0, //0x00000457 movq         %rsi, %rax
	0x48, 0xf7, 0xe2, //0x0000045a mulq         %rdx
	0x48, 0xc1, 0xea, 0x1a, //0x0000045d shrq         $26, %rdx
	0x66, 0x0f, 0x6e, 0xc2, //0x00000461 movd         %edx, %xmm0
	0xf3, 0x0f, 0x6f, 0x0d, 0x93, 0xfb, 0xff, 0xff, //0x00000465 movdqu       $-1133(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	0x66, 0x0f, 0x6f, 0xd8, //0x0000046d movdqa       %xmm0, %xmm3
	0x66, 0x0f, 0xf4, 0xd9, //0x00000471 pmuludq      %xmm1, %xmm3
	0x66, 0x0f, 0x73, 0xd3, 0x2d, //0x00000475 psrlq        $45, %xmm3
	0xb8, 0x10, 0x27, 0x00, 0x00, //0x0000047a movl         $10000, %eax
	0x66, 0x48, 0x0f, 0x6e, 0xd0, //0x0000047f movq         %rax, %xmm2
	0x66, 0x0f, 0x6f, 0xe3, //0x00000484 movdqa       %xmm3, %xmm4
	0x66, 0x0f, 0xf4, 0xe2, //0x00000488 pmuludq      %xmm2, %xmm4
	0x66, 0x0f, 0xfa, 0xc4, //0x0000048c psubd        %xmm4, %xmm0
	0x66, 0x0f, 0x61, 0xd8, //0x00000490 punpcklwd    %xmm0, %xmm3
	0x66, 0x0f, 0x73, 0xf3, 0x02, //0x00000494 psllq        $2, %xmm3
	0xf2, 0x0f, 0x70, 0xc3, 0x50, //0x00000499 pshuflw      $80, %xmm3, %xmm0
	0x66, 0x0f, 0x70, 0xc0, 0x50, //0x0000049e pshufd       $80, %xmm0, %xmm0
	0xf3, 0x0f, 0x6f, 0x25, 0x65, 0xfb, 0xff, 0xff, //0x000004a3 movdqu       $-1179(%rip), %xmm4  /* LCPI0_1+0(%rip) */
	0x66, 0x0f, 0xe4, 0xc4, //0x000004ab pmulhuw      %xmm4, %xmm0
	0xf3, 0x0f, 0x6f, 0x2d, 0x69, 0xfb, 0xff, 0xff, //0x000004af movdqu       $-1175(%rip), %xmm5  /* LCPI0_2+0(%rip) */
	0x66, 0x0f, 0xe4, 0xc5, //0x000004b7 pmulhuw      %xmm5, %xmm0
	0xf3, 0x0f, 0x6f, 0x1d, 0x6d, 0xfb, 0xff, 0xff, //0x000004bb movdqu       $-1171(%rip), %xmm3  /* LCPI0_3+0(%rip) */
	0x66, 0x0f, 0x6f, 0xf0, //0x000004c3 movdqa       %xmm0, %xmm6
	0x66, 0x0f, 0xd5, 0xf3, //0x000004c7 pmullw       %xmm3, %xmm6
	0x66, 0x0f, 0x73, 0xf6, 0x10, //0x000004cb psllq        $16, %xmm6
	0x66, 0x0f, 0xf9, 0xc6, //0x000004d0 psubw        %xmm6, %xmm0
	0x69, 0xc2, 0x00, 0xe1, 0xf5, 0x05, //0x000004d4 imull        $100000000, %edx, %eax
	0x29, 0xc6, //0x000004da subl         %eax, %esi
	0x66, 0x0f, 0x6e, 0xf6, //0x000004dc movd         %esi, %xmm6
	0x66, 0x0f, 0xf4, 0xce, //0x000004e0 pmuludq      %xmm6, %xmm1
	0x66, 0x0f, 0x73, 0xd1, 0x2d, //0x000004e4 psrlq        $45, %xmm1
	0x66, 0x0f, 0xf4, 0xd1, //0x000004e9 pmuludq      %xmm1, %xmm2
	0x66, 0x0f, 0xfa, 0xf2, //0x000004ed psubd        %xmm2, %xmm6
	0x66, 0x0f, 0x61, 0xce, //0x000004f1 punpcklwd    %xmm6, %xmm1
	0x66, 0x0f, 0x73, 0xf1, 0x02, //0x000004f5 psllq        $2, %xmm1
	0xf2, 0x0f, 0x70, 0xc9, 0x50, //0x000004fa pshuflw      $80, %xmm1, %xmm1
	0x66, 0x0f, 0x70, 0xc9, 0x50, //0x000004ff pshufd       $80, %xmm1, %xmm1
	0x66, 0x0f, 0xe4, 0xcc, //0x00000504 pmulhuw      %xmm4, %xmm1
	0x66, 0x0f, 0xe4, 0xcd, //0x00000508 pmulhuw      %xmm5, %xmm1
	0x66, 0x0f, 0xd5, 0xd9, //0x0000050c pmullw       %xmm1, %xmm3
	0x66, 0x0f, 0x73, 0xf3, 0x10, //0x00000510 psllq        $16, %xmm3
	0x66, 0x0f, 0xf9, 0xcb, //0x00000515 psubw        %xmm3, %xmm1
	0x66, 0x0f, 0x67, 0xc1, //0x00000519 packuswb     %xmm1, %xmm0
	0x66, 0x0f, 0xfc, 0x05, 0x1b, 0xfb, 0xff, 0xff, //0x0000051d paddb        $-1253(%rip), %xmm0  /* LCPI0_4+0(%rip) */
	0x89, 0xc8, //0x00000525 movl         %ecx, %eax
	0xf3, 0x0f, 0x7f, 0x04, 0x07, //0x00000527 movdqu       %xmm0, (%rdi,%rax)
	0x83, 0xc9, 0x10, //0x0000052c orl          $16, %ecx
	0x89, 0xc8, //0x0000052f movl         %ecx, %eax
	0x5d, //0x00000531 popq         %rbp
	0xc3, //0x00000532 retq         
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000533 .p2align 4, 0x00
	//0x00000540 _Digits
	0x30, 0x30, 0x30, 0x31, 0x30, 0x32, 0x30, 0x33, 0x30, 0x34, 0x30, 0x35, 0x30, 0x36, 0x30, 0x37, //0x00000540 QUAD $0x3330323031303030; QUAD $0x3730363035303430  // .ascii 16, '0001020304050607'
	0x30, 0x38, 0x30, 0x39, 0x31, 0x30, 0x31, 0x31, 0x31, 0x32, 0x31, 0x33, 0x31, 0x34, 0x31, 0x35, //0x00000550 QUAD $0x3131303139303830; QUAD $0x3531343133313231  // .ascii 16, '0809101112131415'
	0x31, 0x36, 0x31, 0x37, 0x31, 0x38, 0x31, 0x39, 0x32, 0x30, 0x32, 0x31, 0x32, 0x32, 0x32, 0x33, //0x00000560 QUAD $0x3931383137313631; QUAD $0x3332323231323032  // .ascii 16, '1617181920212223'
	0x32, 0x34, 0x32, 0x35, 0x32, 0x36, 0x32, 0x37, 0x32, 0x38, 0x32, 0x39, 0x33, 0x30, 0x33, 0x31, //0x00000570 QUAD $0x3732363235323432; QUAD $0x3133303339323832  // .ascii 16, '2425262728293031'
	0x33, 0x32, 0x33, 0x33, 0x33, 0x34, 0x33, 0x35, 0x33, 0x36, 0x33, 0x37, 0x33, 0x38, 0x33, 0x39, //0x00000580 QUAD $0x3533343333333233; QUAD $0x3933383337333633  // .ascii 16, '3233343536373839'
	0x34, 0x30, 0x34, 0x31, 0x34, 0x32, 0x34, 0x33, 0x34, 0x34, 0x34, 0x35, 0x34, 0x36, 0x34, 0x37, //0x00000590 QUAD $0x3334323431343034; QUAD $0x3734363435343434  // .ascii 16, '4041424344454647'
	0x34, 0x38, 0x34, 0x39, 0x35, 0x30, 0x35, 0x31, 0x35, 0x32, 0x35, 0x33, 0x35, 0x34, 0x35, 0x35, //0x000005a0 QUAD $0x3135303539343834; QUAD $0x3535343533353235  // .ascii 16, '4849505152535455'
	0x35, 0x36, 0x35, 0x37, 0x35, 0x38, 0x35, 0x39, 0x36, 0x30, 0x36, 0x31, 0x36, 0x32, 0x36, 0x33, //0x000005b0 QUAD $0x3935383537353635; QUAD $0x3336323631363036  // .ascii 16, '5657585960616263'
	0x36, 0x34, 0x36, 0x35, 0x36, 0x36, 0x36, 0x37, 0x36, 0x38, 0x36, 0x39, 0x37, 0x30, 0x37, 0x31, //0x000005c0 QUAD $0x3736363635363436; QUAD $0x3137303739363836  // .ascii 16, '6465666768697071'
	0x37, 0x32, 0x37, 0x33, 0x37, 0x34, 0x37, 0x35, 0x37, 0x36, 0x37, 0x37, 0x37, 0x38, 0x37, 0x39, //0x000005d0 QUAD $0x3537343733373237; QUAD $0x3937383737373637  // .ascii 16, '7273747576777879'
	0x38, 0x30, 0x38, 0x31, 0x38, 0x32, 0x38, 0x33, 0x38, 0x34, 0x38, 0x35, 0x38, 0x36, 0x38, 0x37, //0x000005e0 QUAD $0x3338323831383038; QUAD $0x3738363835383438  // .ascii 16, '8081828384858687'
	0x38, 0x38, 0x38, 0x39, 0x39, 0x30, 0x39, 0x31, 0x39, 0x32, 0x39, 0x33, 0x39, 0x34, 0x39, 0x35, //0x000005f0 QUAD $0x3139303939383838; QUAD $0x3539343933393239  // .ascii 16, '8889909192939495'
	0x39, 0x36, 0x39, 0x37, 0x39, 0x38, 0x39, 0x39, //0x00000600 QUAD $0x3939383937393639  // .ascii 8, '96979899'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000608 .p2align 4, 0x00
	//0x00000610 _VecShiftShuffles
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, //0x00000610 QUAD $0x0706050403020100; QUAD $0x0f0e0d0c0b0a0908  // .ascii 16, '\x00\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f'
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, //0x00000620 QUAD $0x0807060504030201; QUAD $0xff0f0e0d0c0b0a09  // .ascii 16, '\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff'
	0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, //0x00000630 QUAD $0x0908070605040302; QUAD $0xffff0f0e0d0c0b0a  // .ascii 16, '\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff'
	0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, //0x00000640 QUAD $0x0a09080706050403; QUAD $0xffffff0f0e0d0c0b  // .ascii 16, '\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff'
	0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, //0x00000650 QUAD $0x0b0a090807060504; QUAD $0xffffffff0f0e0d0c  // .ascii 16, '\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff'
	0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, //0x00000660 QUAD $0x0c0b0a0908070605; QUAD $0xffffffffff0f0e0d  // .ascii 16, '\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff'
	0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, //0x00000670 QUAD $0x0d0c0b0a09080706; QUAD $0xffffffffffff0f0e  // .ascii 16, '\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff'
	0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, //0x00000680 QUAD $0x0e0d0c0b0a090807; QUAD $0xffffffffffffff0f  // .ascii 16, '\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff\xff'
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, //0x00000690 QUAD $0x0f0e0d0c0b0a0908; QUAD $0xffffffffffffffff  // .ascii 16, '\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff'
}
 