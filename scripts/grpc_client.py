import asyncio

from grpclib.client import Channel

from ozonmp.omp_Product_api.v1.omp_Product_api_grpc import GoProductApiServiceStub
from ozonmp.omp_Product_api.v1.omp_Product_api_pb2 import DescribeProductV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = GoProductApiServiceStub(channel)

        req = DescribeProductV1Request(Product_id=1)
        reply = await client.DescribeProductV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
