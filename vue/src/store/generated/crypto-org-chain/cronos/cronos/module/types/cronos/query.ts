/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'cronos'

/** ContractByDenomRequest is the request type of ContractByDenom call */
export interface ContractByDenomRequest {
  denom: string
}

/** ContractByDenomRequest is the response type of ContractByDenom call */
export interface ContractByDenomResponse {
  contract: string
  autoContract: string
}

/** DenomByContractRequest is the request type of DenomByContract call */
export interface DenomByContractRequest {
  contract: string
}

/** DenomByContractResponse is the response type of DenomByContract call */
export interface DenomByContractResponse {
  denom: string
}

export interface QueryInterchainAccountRequest {
  connectionId: string
  owner: string
}

export interface QueryInterchainAccountResponse {
  interchainAccountAddress: string
}

const baseContractByDenomRequest: object = { denom: '' }

export const ContractByDenomRequest = {
  encode(message: ContractByDenomRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ContractByDenomRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseContractByDenomRequest } as ContractByDenomRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ContractByDenomRequest {
    const message = { ...baseContractByDenomRequest } as ContractByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: ContractByDenomRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<ContractByDenomRequest>): ContractByDenomRequest {
    const message = { ...baseContractByDenomRequest } as ContractByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseContractByDenomResponse: object = { contract: '', autoContract: '' }

export const ContractByDenomResponse = {
  encode(message: ContractByDenomResponse, writer: Writer = Writer.create()): Writer {
    if (message.contract !== '') {
      writer.uint32(10).string(message.contract)
    }
    if (message.autoContract !== '') {
      writer.uint32(18).string(message.autoContract)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ContractByDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseContractByDenomResponse } as ContractByDenomResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.contract = reader.string()
          break
        case 2:
          message.autoContract = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ContractByDenomResponse {
    const message = { ...baseContractByDenomResponse } as ContractByDenomResponse
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = String(object.contract)
    } else {
      message.contract = ''
    }
    if (object.autoContract !== undefined && object.autoContract !== null) {
      message.autoContract = String(object.autoContract)
    } else {
      message.autoContract = ''
    }
    return message
  },

  toJSON(message: ContractByDenomResponse): unknown {
    const obj: any = {}
    message.contract !== undefined && (obj.contract = message.contract)
    message.autoContract !== undefined && (obj.autoContract = message.autoContract)
    return obj
  },

  fromPartial(object: DeepPartial<ContractByDenomResponse>): ContractByDenomResponse {
    const message = { ...baseContractByDenomResponse } as ContractByDenomResponse
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = object.contract
    } else {
      message.contract = ''
    }
    if (object.autoContract !== undefined && object.autoContract !== null) {
      message.autoContract = object.autoContract
    } else {
      message.autoContract = ''
    }
    return message
  }
}

const baseDenomByContractRequest: object = { contract: '' }

export const DenomByContractRequest = {
  encode(message: DenomByContractRequest, writer: Writer = Writer.create()): Writer {
    if (message.contract !== '') {
      writer.uint32(10).string(message.contract)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): DenomByContractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseDenomByContractRequest } as DenomByContractRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.contract = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): DenomByContractRequest {
    const message = { ...baseDenomByContractRequest } as DenomByContractRequest
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = String(object.contract)
    } else {
      message.contract = ''
    }
    return message
  },

  toJSON(message: DenomByContractRequest): unknown {
    const obj: any = {}
    message.contract !== undefined && (obj.contract = message.contract)
    return obj
  },

  fromPartial(object: DeepPartial<DenomByContractRequest>): DenomByContractRequest {
    const message = { ...baseDenomByContractRequest } as DenomByContractRequest
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = object.contract
    } else {
      message.contract = ''
    }
    return message
  }
}

const baseDenomByContractResponse: object = { denom: '' }

export const DenomByContractResponse = {
  encode(message: DenomByContractResponse, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): DenomByContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseDenomByContractResponse } as DenomByContractResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): DenomByContractResponse {
    const message = { ...baseDenomByContractResponse } as DenomByContractResponse
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: DenomByContractResponse): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<DenomByContractResponse>): DenomByContractResponse {
    const message = { ...baseDenomByContractResponse } as DenomByContractResponse
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryInterchainAccountRequest: object = { connectionId: '', owner: '' }

export const QueryInterchainAccountRequest = {
  encode(message: QueryInterchainAccountRequest, writer: Writer = Writer.create()): Writer {
    if (message.connectionId !== '') {
      writer.uint32(10).string(message.connectionId)
    }
    if (message.owner !== '') {
      writer.uint32(18).string(message.owner)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryInterchainAccountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryInterchainAccountRequest } as QueryInterchainAccountRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.connectionId = reader.string()
          break
        case 2:
          message.owner = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryInterchainAccountRequest {
    const message = { ...baseQueryInterchainAccountRequest } as QueryInterchainAccountRequest
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = String(object.connectionId)
    } else {
      message.connectionId = ''
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    return message
  },

  toJSON(message: QueryInterchainAccountRequest): unknown {
    const obj: any = {}
    message.connectionId !== undefined && (obj.connectionId = message.connectionId)
    message.owner !== undefined && (obj.owner = message.owner)
    return obj
  },

  fromPartial(object: DeepPartial<QueryInterchainAccountRequest>): QueryInterchainAccountRequest {
    const message = { ...baseQueryInterchainAccountRequest } as QueryInterchainAccountRequest
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = object.connectionId
    } else {
      message.connectionId = ''
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
    }
    return message
  }
}

const baseQueryInterchainAccountResponse: object = { interchainAccountAddress: '' }

export const QueryInterchainAccountResponse = {
  encode(message: QueryInterchainAccountResponse, writer: Writer = Writer.create()): Writer {
    if (message.interchainAccountAddress !== '') {
      writer.uint32(10).string(message.interchainAccountAddress)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryInterchainAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryInterchainAccountResponse } as QueryInterchainAccountResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.interchainAccountAddress = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryInterchainAccountResponse {
    const message = { ...baseQueryInterchainAccountResponse } as QueryInterchainAccountResponse
    if (object.interchainAccountAddress !== undefined && object.interchainAccountAddress !== null) {
      message.interchainAccountAddress = String(object.interchainAccountAddress)
    } else {
      message.interchainAccountAddress = ''
    }
    return message
  },

  toJSON(message: QueryInterchainAccountResponse): unknown {
    const obj: any = {}
    message.interchainAccountAddress !== undefined && (obj.interchainAccountAddress = message.interchainAccountAddress)
    return obj
  },

  fromPartial(object: DeepPartial<QueryInterchainAccountResponse>): QueryInterchainAccountResponse {
    const message = { ...baseQueryInterchainAccountResponse } as QueryInterchainAccountResponse
    if (object.interchainAccountAddress !== undefined && object.interchainAccountAddress !== null) {
      message.interchainAccountAddress = object.interchainAccountAddress
    } else {
      message.interchainAccountAddress = ''
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** ContractByDenom queries contract addresses by native denom */
  ContractByDenom(request: ContractByDenomRequest): Promise<ContractByDenomResponse>
  /** DenomByContract queries native denom by contract address */
  DenomByContract(request: DenomByContractRequest): Promise<DenomByContractResponse>
  /** Queries a list of InterchainAccount items. */
  InterchainAccount(request: QueryInterchainAccountRequest): Promise<QueryInterchainAccountResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  ContractByDenom(request: ContractByDenomRequest): Promise<ContractByDenomResponse> {
    const data = ContractByDenomRequest.encode(request).finish()
    const promise = this.rpc.request('cronos.Query', 'ContractByDenom', data)
    return promise.then((data) => ContractByDenomResponse.decode(new Reader(data)))
  }

  DenomByContract(request: DenomByContractRequest): Promise<DenomByContractResponse> {
    const data = DenomByContractRequest.encode(request).finish()
    const promise = this.rpc.request('cronos.Query', 'DenomByContract', data)
    return promise.then((data) => DenomByContractResponse.decode(new Reader(data)))
  }

  InterchainAccount(request: QueryInterchainAccountRequest): Promise<QueryInterchainAccountResponse> {
    const data = QueryInterchainAccountRequest.encode(request).finish()
    const promise = this.rpc.request('cronos.Query', 'InterchainAccount', data)
    return promise.then((data) => QueryInterchainAccountResponse.decode(new Reader(data)))
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
