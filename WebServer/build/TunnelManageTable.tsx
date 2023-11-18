// ** React Imports
import { useEffect, useState } from 'react'

// ** MUI Imports
import Paper from '@mui/material/Paper'
import Table from '@mui/material/Table'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import TableHead from '@mui/material/TableHead'
import TableRow from '@mui/material/TableRow'

// ** Icons Imports
import { Box, Button, Card, CardHeader, Typography } from '@mui/material'
import { useStore } from 'reto'
import { Tunnel } from 'src/service/models'
import { getPolicies } from 'src/service/policyApi'
import { deleteTunnel, getTunnelDetailInfo, getTunnels, updateTunnel } from 'src/service/tunnelApi'
import { RightDrawerStore } from 'src/states/global/RightDrawerStore'
import { FileSharePageStore } from 'src/states/page/file-share'
import { PolicySelectData, TunnelDataInfos } from 'src/types/models/DataModels'
import { SnackBarStore } from '../../states/global/TopSnackBarStore'
import { ConfimDeleteDialog } from '../common/ComfirmDeleteDialog'
import { DetailContent } from '../common/DataDetailContent'
import Loading from '../common/LoadingProgress'
import { ModalDialogWithoutButton } from '../common/ModalDialog'
import { TableBodyRow } from '../common/Tables/CollapseTableRow'
import Toast from '../common/Toast'
import { TunnelCreateTabs } from './TunnelCreateForm'
import { TunnelUpdateTabs } from './TunnelUpdateForm'

const TunnelManageTable = () => {
  //页面所需状态
  const states = useStore(FileSharePageStore)
  const topSnackBarStates = useStore(SnackBarStore)
  const rightDrawerStates = useStore(RightDrawerStore)
  //当前需要删除的通道id

  const [needDeleteId, setNeedDeleteId] = useState<number>(-1)
  const [needUpdateId, setNeedUpdateId] = useState<number>(-1)
  const [dataInfos, setDataInfos] = useState<Array<TunnelDataInfos>>([]);
  const { rows, loading } = states
  //加载通道数据
  useEffect(() => {
    getAllTunnel();
  }, [])

  //通道表头
  const subTableFields = ['id', '名称', '唯一id', '提供目标', '创建时间']

  // //计算通道子表数据
  // const getSubTableValues = (row: TunnelData) => row.subTableData.map(subTableRow => {
  //   return [`${subTableRow.id}`, subTableRow.name, subTableRow.uniqueId, `${subTableRow.target}`, `${toLocalTimeString(subTableRow.createTime)}`]
  // })


  //计算通道表数据
  // const getRowValues = (row: TunnelData) => {
  //   return [`${row.id}`, `${row.name}`, `${row.policyName}`, `${row.owner}`, `${row.provideDataCount}`, `${toLocalTimeString(row.createTime)}`]
  // }

  //只传了策略id 作修改
  const getRowValues = (row: Tunnel) => {
    return [`${row.id}`, `${row.name}`, `${row.policyId}`, `${row.ownerId}`, `${row.createdTime}`, `${row.updatedTime}`]
  }



  const handleAddTunnel = () => {
    //新建数据通道时，需请求现有策略
    getPolicies().then(res => {
      if (res) {
        console.log('policys')
        console.log(res)
        let policies
        if (Array.isArray(res.data)) {
          policies = res.data
        } else {
          policies = [res.data]
        }
        states.setPolicySelectData(policies.map((p, index) => { return { id: p.id, name: p.name, index: index } }))
      } else {

      }
    })
    states.setEditorDialogOpen(true)
  }

  const handleDeleteClick = (id: number) => {
    setNeedDeleteId(id)
    states.setConfirmDeleteDialogOpen(true)
  }
  const handleDetailClick = (id: number) => {
    rightDrawerStates.setRightDrawerOpen(true);
    getTunnelDetailInfo(id).then(res => {
      console.log("获取数据通道的详情")
      console.log(res)
      // setDataInfos(res?.data.dataInfos);
      // 这里本来想直接组件状态的，但是是一个异步操作
      rightDrawerStates.setRightDrawerContent(
        <DetailContent title={'数据详情'} items={res?.data.dataInfos} />
      )
    })
  }

  //更新特定id的数据通道
  const handleUpdateClick = async (id: number) => {
    //更新数据通道时，需请求现有策略
    await getPolicies().then(res => {
      let policies
      if (Array.isArray(res.data)) {
        policies = res.data
      } else {
        policies = [res.data]
      }
      states.setPolicySelectData(
        policies.map((policy, index) => {
          let policySelectData: PolicySelectData = {
            index: index,
            id: policy.id,
            name: policy.name
          }
          return policySelectData
        })
      )
    }).catch(err => {
      topSnackBarStates.setSnackBarOpen(true)
      topSnackBarStates.setSnackBarType('error')
      topSnackBarStates.setSnackBarMessage(err)
    })
    await getTunnelDetailInfo(id).then(res => {
      if (res) {
        console.log('更新数据通道，获取现有信息')
        console.log(res)
        let data = res.data;
        states.setTunnelName(data.name);
        states.setTunnelDes(data.description);
        states.setPolicyId(data.policy.id);

      }
    })
    setNeedUpdateId(id);
    states.setUpdateDialogOpen(true)
  }

  //更新数据通道
  const updateTunnelLogic = (data: object) => {
    updateTunnel(data).then(res => {
      console.log("更新数据通道");
      console.log(res);
      if (res?.status == 0) {
        getAllTunnel();
        Toast.success("更新成功")
      } else {
        Toast.error("请求失败");
      }
    }).catch(e => {
      Toast.error("请求失败");
    })
    states.setUpdateDialogOpen(false)
  }

  //删除特定id的数据通道
  const deleteTunnelLogic = (id: number) => {
    deleteTunnel(id).then(res => {
      console.log("删除数据通道");
      console.log(res);
      if (res?.status == 0) {
        getAllTunnel();
        Toast.success("删除成功")
      } else {
        Toast.error("请求失败");
      }
    }).catch(e => {
      Toast.error("请求失败");
    })
  }

  //删除特定id的数据通道
  const deleteTunnelByIdLogic = (id: number) => {
    deleteTunnel(id)
      .then(res => {
        console.log('删除数据通道')
        console.log(res)
        if (res?.status == 0) {
          getAllTunnel()
          Toast.success('删除成功')
        } else {
          Toast.error('请求失败')
        }
      })
      .catch(e => {
        Toast.error('请求失败')
      })
  }

  //获取所有数据通道信息
  const getAllTunnel = () => {
    getTunnels().then(res => {
      console.log('useEffect')
      console.log(res)
      // if (res && res.code == 200) {
      if (res) {
        if (Array.isArray(res.data)) {
          states.setRows(res.data)
        } else {
          states.setRows([res.data])
        }
        states.setLoading(false)
      } else {
        topSnackBarStates.setSnackBarMessage('加载失败')
        topSnackBarStates.setSnackBarOpen(true)
        topSnackBarStates.setSnackBarType('error')
      }
    })
  }

  return (
    // 加载组件
    <Loading loading={loading}>
      {/* 共用的对话框、抽屉、浮动窗口组件 */}
      <ConfimDeleteDialog
        open={states.confirmDeleteDialogOpen}
        onCancelClick={() => states.setConfirmDeleteDialogOpen(false)}
        onOkClick={() => {
          states.setConfirmDeleteDialogOpen(false)
          deleteTunnelLogic(needDeleteId)
          console.log(`remove:${needDeleteId}`)
        }}
      />
      <ModalDialogWithoutButton maxWidth='lg' open={states.editorDialogOpen} title={'新建数据通道'}>
        <TunnelCreateTabs getAll={getAllTunnel} />
      </ModalDialogWithoutButton>
      <ModalDialogWithoutButton maxWidth='lg' open={states.updateDialogOpen} title={'更新数据通道'}>
        <TunnelUpdateTabs id={needUpdateId} updateById={updateTunnelLogic} />
      </ModalDialogWithoutButton>
      {/* 表内容 */}
      <Card>
        <CardHeader
          title={
            <Box display='flex' justifyContent='space-between'>
              <Typography variant='h6'>数据通道概览</Typography>
              <Button variant='contained' onClick={handleAddTunnel}>
                新建数据通道
              </Button>
            </Box>
          }
          titleTypographyProps={{ variant: 'h6' }}
        />
        <TableContainer component={Paper}>
          <Table aria-label='collapsible table'>
            <TableHead>
              <TableRow>
                <TableCell align='center' />
                <TableCell align='center'>ID</TableCell>
                <TableCell align='center' sx={{ minWidth: 200 }}>
                  名称
                </TableCell>
                <TableCell align='center' sx={{ minWidth: 300 }}>
                  使用策略ID
                </TableCell>
                <TableCell align='center'>拥有者ID</TableCell>
                {/* <TableCell align='center'>传输数据次数</TableCell> */}
                <TableCell align='center'>创建时间</TableCell>
                <TableCell align='center'>更新时间</TableCell>
                <TableCell align='center' sx={{ minWidth: 400 }}>
                  操作
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {/* {rows?.map(row => (
                  <CollapseTableRow
                    key={row.id}
                    values={getRowValues(row)}
                    subTableProps={{ title: '数据传输历史', fields: subTableFields, values: getSubTableValues(row) }}
                    operations={
                      <>
                        <Button
                          onClick={() => handleDetailClick(row.id)}
                          sx={{ ml: 2, mr: 2 }}
                          style={{ color: 'white' }}
                          variant="contained"
                        >详情</Button>
                        <Button
                          onClick={() => handleDeleteClick(row.id)}
                          variant="contained"
                          style={{ color: 'white' }}
                          sx={{ ml: 2, mr: 2 }}
                          color='error'
                        >删除</Button>
                      </>
                    }
                  />
                ))} */}
              {rows?.map(row => (
                <TableBodyRow
                  key={row.id}
                  values={getRowValues(row)}
                  operations={
                    <>
                      <Button
                        onClick={() => handleDetailClick(row.id)}
                        sx={{ ml: 1, mr: 1 }}
                        style={{ color: 'white' }}
                        variant="contained"
                      >详情</Button>
                      <Button
                        onClick={() => handleUpdateClick(row.id)}
                        variant="contained"
                        style={{ color: 'white' }}
                        sx={{ ml: 1, mr: 1 }}
                        color='success'
                      >更新</Button>
                      <Button
                        onClick={() => handleDeleteClick(row.id)}
                        variant="contained"
                        style={{ color: 'white' }}
                        sx={{ ml: 1, mr: 1 }}
                        color='error'
                      >删除</Button>
                    </>
                  }
                />
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>
    </Loading>
  )
}

export default TunnelManageTable
