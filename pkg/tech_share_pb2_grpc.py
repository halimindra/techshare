# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import tech_share_pb2 as tech__share__pb2


class TechShareStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetPerson = channel.unary_unary(
        '/pkg.TechShare/GetPerson',
        request_serializer=tech__share__pb2.PersonRequest.SerializeToString,
        response_deserializer=tech__share__pb2.Person.FromString,
        )
    self.ListPeople = channel.unary_stream(
        '/pkg.TechShare/ListPeople',
        request_serializer=tech__share__pb2.PeopleRequest.SerializeToString,
        response_deserializer=tech__share__pb2.Person.FromString,
        )


class TechShareServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetPerson(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListPeople(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TechShareServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetPerson': grpc.unary_unary_rpc_method_handler(
          servicer.GetPerson,
          request_deserializer=tech__share__pb2.PersonRequest.FromString,
          response_serializer=tech__share__pb2.Person.SerializeToString,
      ),
      'ListPeople': grpc.unary_stream_rpc_method_handler(
          servicer.ListPeople,
          request_deserializer=tech__share__pb2.PeopleRequest.FromString,
          response_serializer=tech__share__pb2.Person.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'pkg.TechShare', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
