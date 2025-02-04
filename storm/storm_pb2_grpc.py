# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import storm_pb2 as storm__pb2


class StormServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Summarize = channel.unary_unary(
        '/StormService/Summarize',
        request_serializer=storm__pb2.SummarizeRequest.SerializeToString,
        response_deserializer=storm__pb2.SummarizeResponse.FromString,
        )
    self.HadistRetrieval = channel.unary_unary(
        '/StormService/HadistRetrieval',
        request_serializer=storm__pb2.HadistRequest.SerializeToString,
        response_deserializer=storm__pb2.HadistReponse.FromString,
        )


class StormServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Summarize(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def HadistRetrieval(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_StormServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Summarize': grpc.unary_unary_rpc_method_handler(
          servicer.Summarize,
          request_deserializer=storm__pb2.SummarizeRequest.FromString,
          response_serializer=storm__pb2.SummarizeResponse.SerializeToString,
      ),
      'HadistRetrieval': grpc.unary_unary_rpc_method_handler(
          servicer.HadistRetrieval,
          request_deserializer=storm__pb2.HadistRequest.FromString,
          response_serializer=storm__pb2.HadistReponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'StormService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
