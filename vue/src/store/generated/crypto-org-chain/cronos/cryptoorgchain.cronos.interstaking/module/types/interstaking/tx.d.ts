import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "cryptoorgchain.cronos.interstaking";
export interface MsgRegisterAccount {
    owner: string;
    connectionId: string;
}
export interface MsgRegisterAccountResponse {
}
export declare const MsgRegisterAccount: {
    encode(message: MsgRegisterAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterAccount;
    fromJSON(object: any): MsgRegisterAccount;
    toJSON(message: MsgRegisterAccount): unknown;
    fromPartial(object: DeepPartial<MsgRegisterAccount>): MsgRegisterAccount;
};
export declare const MsgRegisterAccountResponse: {
    encode(_: MsgRegisterAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterAccountResponse;
    fromJSON(_: any): MsgRegisterAccountResponse;
    toJSON(_: MsgRegisterAccountResponse): unknown;
    fromPartial(_: DeepPartial<MsgRegisterAccountResponse>): MsgRegisterAccountResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RegisterAccount(request: MsgRegisterAccount): Promise<MsgRegisterAccountResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    RegisterAccount(request: MsgRegisterAccount): Promise<MsgRegisterAccountResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
