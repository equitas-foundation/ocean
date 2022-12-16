# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from ocean.v1 import account_pb2 as ocean_dot_v1_dot_account__pb2


class AccountServiceStub(object):
    """AccountService is used to manage accounts in HD Wallet.
    It supports generating addresses, listing utxos and balances for specific account or
    list of addresses and selecting utxos.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateAccountBIP44 = channel.unary_unary(
                '/ocean.v1.AccountService/CreateAccountBIP44',
                request_serializer=ocean_dot_v1_dot_account__pb2.CreateAccountBIP44Request.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.CreateAccountBIP44Response.FromString,
                )
        self.CreateAccountMultiSig = channel.unary_unary(
                '/ocean.v1.AccountService/CreateAccountMultiSig',
                request_serializer=ocean_dot_v1_dot_account__pb2.CreateAccountMultiSigRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.CreateAccountMultiSigResponse.FromString,
                )
        self.CreateAccountCustom = channel.unary_unary(
                '/ocean.v1.AccountService/CreateAccountCustom',
                request_serializer=ocean_dot_v1_dot_account__pb2.CreateAccountCustomRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.CreateAccountCustomResponse.FromString,
                )
        self.SetAccountTemplate = channel.unary_unary(
                '/ocean.v1.AccountService/SetAccountTemplate',
                request_serializer=ocean_dot_v1_dot_account__pb2.SetAccountTemplateRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.SetAccountTemplateResponse.FromString,
                )
        self.DeriveAddresses = channel.unary_unary(
                '/ocean.v1.AccountService/DeriveAddresses',
                request_serializer=ocean_dot_v1_dot_account__pb2.DeriveAddressesRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.DeriveAddressesResponse.FromString,
                )
        self.DeriveChangeAddresses = channel.unary_unary(
                '/ocean.v1.AccountService/DeriveChangeAddresses',
                request_serializer=ocean_dot_v1_dot_account__pb2.DeriveChangeAddressesRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.DeriveChangeAddressesResponse.FromString,
                )
        self.ListAddresses = channel.unary_unary(
                '/ocean.v1.AccountService/ListAddresses',
                request_serializer=ocean_dot_v1_dot_account__pb2.ListAddressesRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.ListAddressesResponse.FromString,
                )
        self.Balance = channel.unary_unary(
                '/ocean.v1.AccountService/Balance',
                request_serializer=ocean_dot_v1_dot_account__pb2.BalanceRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.BalanceResponse.FromString,
                )
        self.ListUtxos = channel.unary_unary(
                '/ocean.v1.AccountService/ListUtxos',
                request_serializer=ocean_dot_v1_dot_account__pb2.ListUtxosRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.ListUtxosResponse.FromString,
                )
        self.DeleteAccount = channel.unary_unary(
                '/ocean.v1.AccountService/DeleteAccount',
                request_serializer=ocean_dot_v1_dot_account__pb2.DeleteAccountRequest.SerializeToString,
                response_deserializer=ocean_dot_v1_dot_account__pb2.DeleteAccountResponse.FromString,
                )


class AccountServiceServicer(object):
    """AccountService is used to manage accounts in HD Wallet.
    It supports generating addresses, listing utxos and balances for specific account or
    list of addresses and selecting utxos.
    """

    def CreateAccountBIP44(self, request, context):
        """CreateAccountBIP44 creates a new BIP44 account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateAccountMultiSig(self, request, context):
        """CreateAccountMultiSig creates a new multisig account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateAccountCustom(self, request, context):
        """CreateAccountCustom creates a new custom account for which loading a template.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetAccountTemplate(self, request, context):
        """SetAccountTemplate sets the template for the account used to generate new addresses.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeriveAddresses(self, request, context):
        """DeriveAddresses generates new address(es) for the account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeriveChangeAddresses(self, request, context):
        """DeriveChangeAddresses generates new change address(es) for the account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListAddresses(self, request, context):
        """ListAddresses returns all derived addresses for the account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Balance(self, request, context):
        """Balance returns the balance for the account, or for specific list of 
        account's addresses.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListUtxos(self, request, context):
        """ListUtxos returns the utxos for the account, or specific list of 
        account's addresses.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteAccount(self, request, context):
        """DeleteAccount deletes an existing account. The operation is allowed only
        if the account has zero balance.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AccountServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateAccountBIP44': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateAccountBIP44,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.CreateAccountBIP44Request.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.CreateAccountBIP44Response.SerializeToString,
            ),
            'CreateAccountMultiSig': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateAccountMultiSig,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.CreateAccountMultiSigRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.CreateAccountMultiSigResponse.SerializeToString,
            ),
            'CreateAccountCustom': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateAccountCustom,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.CreateAccountCustomRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.CreateAccountCustomResponse.SerializeToString,
            ),
            'SetAccountTemplate': grpc.unary_unary_rpc_method_handler(
                    servicer.SetAccountTemplate,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.SetAccountTemplateRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.SetAccountTemplateResponse.SerializeToString,
            ),
            'DeriveAddresses': grpc.unary_unary_rpc_method_handler(
                    servicer.DeriveAddresses,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.DeriveAddressesRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.DeriveAddressesResponse.SerializeToString,
            ),
            'DeriveChangeAddresses': grpc.unary_unary_rpc_method_handler(
                    servicer.DeriveChangeAddresses,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.DeriveChangeAddressesRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.DeriveChangeAddressesResponse.SerializeToString,
            ),
            'ListAddresses': grpc.unary_unary_rpc_method_handler(
                    servicer.ListAddresses,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.ListAddressesRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.ListAddressesResponse.SerializeToString,
            ),
            'Balance': grpc.unary_unary_rpc_method_handler(
                    servicer.Balance,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.BalanceRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.BalanceResponse.SerializeToString,
            ),
            'ListUtxos': grpc.unary_unary_rpc_method_handler(
                    servicer.ListUtxos,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.ListUtxosRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.ListUtxosResponse.SerializeToString,
            ),
            'DeleteAccount': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteAccount,
                    request_deserializer=ocean_dot_v1_dot_account__pb2.DeleteAccountRequest.FromString,
                    response_serializer=ocean_dot_v1_dot_account__pb2.DeleteAccountResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ocean.v1.AccountService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AccountService(object):
    """AccountService is used to manage accounts in HD Wallet.
    It supports generating addresses, listing utxos and balances for specific account or
    list of addresses and selecting utxos.
    """

    @staticmethod
    def CreateAccountBIP44(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/CreateAccountBIP44',
            ocean_dot_v1_dot_account__pb2.CreateAccountBIP44Request.SerializeToString,
            ocean_dot_v1_dot_account__pb2.CreateAccountBIP44Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateAccountMultiSig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/CreateAccountMultiSig',
            ocean_dot_v1_dot_account__pb2.CreateAccountMultiSigRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.CreateAccountMultiSigResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateAccountCustom(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/CreateAccountCustom',
            ocean_dot_v1_dot_account__pb2.CreateAccountCustomRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.CreateAccountCustomResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetAccountTemplate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/SetAccountTemplate',
            ocean_dot_v1_dot_account__pb2.SetAccountTemplateRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.SetAccountTemplateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeriveAddresses(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/DeriveAddresses',
            ocean_dot_v1_dot_account__pb2.DeriveAddressesRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.DeriveAddressesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeriveChangeAddresses(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/DeriveChangeAddresses',
            ocean_dot_v1_dot_account__pb2.DeriveChangeAddressesRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.DeriveChangeAddressesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListAddresses(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/ListAddresses',
            ocean_dot_v1_dot_account__pb2.ListAddressesRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.ListAddressesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Balance(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/Balance',
            ocean_dot_v1_dot_account__pb2.BalanceRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.BalanceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListUtxos(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/ListUtxos',
            ocean_dot_v1_dot_account__pb2.ListUtxosRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.ListUtxosResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteAccount(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1.AccountService/DeleteAccount',
            ocean_dot_v1_dot_account__pb2.DeleteAccountRequest.SerializeToString,
            ocean_dot_v1_dot_account__pb2.DeleteAccountResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
