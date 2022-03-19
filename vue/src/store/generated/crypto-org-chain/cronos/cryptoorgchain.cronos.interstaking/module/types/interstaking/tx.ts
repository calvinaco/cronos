/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'cryptoorgchain.cronos.interstaking'

export interface MsgRegisterAccount {
  owner: string
  connectionId: string
}

export interface MsgRegisterAccountResponse {}

const baseMsgRegisterAccount: object = { owner: '', connectionId: '' }

export const MsgRegisterAccount = {
  encode(message: MsgRegisterAccount, writer: Writer = Writer.create()): Writer {
    if (message.owner !== '') {
      writer.uint32(10).string(message.owner)
    }
    if (message.connectionId !== '') {
      writer.uint32(18).string(message.connectionId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRegisterAccount } as MsgRegisterAccount
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string()
          break
        case 2:
          message.connectionId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgRegisterAccount {
    const message = { ...baseMsgRegisterAccount } as MsgRegisterAccount
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = String(object.connectionId)
    } else {
      message.connectionId = ''
    }
    return message
  },

  toJSON(message: MsgRegisterAccount): unknown {
    const obj: any = {}
    message.owner !== undefined && (obj.owner = message.owner)
    message.connectionId !== undefined && (obj.connectionId = message.connectionId)
    return obj
  },

  fromPartial(object: DeepPartial<MsgRegisterAccount>): MsgRegisterAccount {
    const message = { ...baseMsgRegisterAccount } as MsgRegisterAccount
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
    }
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = object.connectionId
    } else {
      message.connectionId = ''
    }
    return message
  }
}

const baseMsgRegisterAccountResponse: object = {}

export const MsgRegisterAccountResponse = {
  encode(_: MsgRegisterAccountResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRegisterAccountResponse } as MsgRegisterAccountResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgRegisterAccountResponse {
    const message = { ...baseMsgRegisterAccountResponse } as MsgRegisterAccountResponse
    return message
  },

  toJSON(_: MsgRegisterAccountResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgRegisterAccountResponse>): MsgRegisterAccountResponse {
    const message = { ...baseMsgRegisterAccountResponse } as MsgRegisterAccountResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RegisterAccount(request: MsgRegisterAccount): Promise<MsgRegisterAccountResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  RegisterAccount(request: MsgRegisterAccount): Promise<MsgRegisterAccountResponse> {
    const data = MsgRegisterAccount.encode(request).finish()
    const promise = this.rpc.request('cryptoorgchain.cronos.interstaking.Msg', 'RegisterAccount', data)
    return promise.then((data) => MsgRegisterAccountResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
