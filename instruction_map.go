package sm83

var InstructionMap = map[byte]Instruction{
	0x00: {IK: IK_NOP, AM: AM_IMP},
	0x01: {IK: IK_LD, AM: AM_R_D16, R1: RK_BC},
	0x02: {IK: IK_LD, AM: AM_MR_R, R1: RK_BC, R2: RK_A},
	0x11: {IK: IK_LD, AM: AM_R_D16, R1: RK_DE},
	0x21: {IK: IK_LD, AM: AM_R_D16, R1: RK_HL},
	0x31: {IK: IK_LD, AM: AM_R_D16, R1: RK_SP},
	0x40: {IK: IK_NOP, AM: AM_IMP},
	0x49: {IK: IK_NOP, AM: AM_IMP},
	0x52: {IK: IK_NOP, AM: AM_IMP},
	0x5b: {IK: IK_NOP, AM: AM_IMP},
	0x64: {IK: IK_NOP, AM: AM_IMP},
	0x6d: {IK: IK_NOP, AM: AM_IMP},
	0x7f: {IK: IK_NOP, AM: AM_IMP},
	0xC2: {IK: IK_JP, AM: AM_D16, CK: CK_NZ},
	0xC3: {IK: IK_JP, AM: AM_D16, CK: CK_NONE},
	0xCA: {IK: IK_JP, AM: AM_D16, CK: CK_Z},
	0xD2: {IK: IK_JP, AM: AM_D16, CK: CK_NC},
	0xDA: {IK: IK_JP, AM: AM_D16, CK: CK_C},
	0xE9: {IK: IK_JP, AM: AM_R, R1: RK_HL, CK: CK_NONE},
}
