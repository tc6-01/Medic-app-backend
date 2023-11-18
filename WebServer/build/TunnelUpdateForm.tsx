import { Box, Button, Grid, Stack, Step, StepButton, Stepper, Typography, Alert } from '@mui/material'
import * as React from 'react'
import { useStore } from 'reto'
import { TunnelCreateDialogStore, FileSharePageStore } from 'src/states/page/file-share'
import { SelectInput, TextInput } from '../common/InputControls'
import { addTunnel } from 'src/service/tunnelApi'
import Toast from '../common/Toast'
import { UpdateProps } from 'src/types/ComponentProps'

const BasicInfoForm = () => {
  const states = useStore(TunnelCreateDialogStore);
  return (
    <Stack spacing={5} p={3} alignItems='space-around'>
      <TextInput name={'名称'} placeHolder={'请输入名称'} defaultValue={states.tunnelName} onChange={(e) => { states.setTunnelName(e.target.value) }} />
      <TextInput name={'数据通道详情'} placeHolder={'请输入描述'} defaultValue={states.tunnelDes} onChange={(e) => { states.setTunnelDes(e.target.value) }} />
      <SelectInput name={'绑定策略'} options={states.policySelectData} defaultValue={states.policyId} placeHolder={'选择要绑定的策略'} onChange={(e) => { states.setPolicyId(e.target.value) }} />
      {/* <SelectInput name={'拥有者'} options={['拥有者1']} placeHolder={'选择要绑定的拥有者'} onChange={onChange}/> */}
    </Stack>
  )
}

const TunnelUpdateTabs = ({ id, updateById }: UpdateProps) => {
  // 状态一样，复用TunnelCrate的状态
  const states = useStore(TunnelCreateDialogStore)
  const dialogStates = useStore(FileSharePageStore)


  //取消弹框
  const handleCancel = () => {
    dialogStates.setUpdateDialogOpen(false);
  }

  //更新数据通道请求
  const onSubmit = async () => {
    let data = { id: id, Name: states.tunnelName, Description: states.tunnelDes, PolicyId: states.policyId }
    let res = await updateById(data)
    console.log('更新数据通道')
    console.log(res)
  }

  return (
    <Grid sx={{ minHeight: 500, height: '100%', pl: 10, pr: 10 }} >
      <Stack sx={{ height: 400 }}>
        {/* <Stepper nonLinear activeStep={states.activeStep}>
                    {steps.map((label, index) => (
                        <Step key={label} completed={states.completed[index]}>
                            <StepButton color="inherit" onClick={handleStep(index)}>
                                {label}
                            </StepButton>
                        </Step>
                    ))}
                </Stepper> */}
        <Grid container xs={12} sx={{ pt: 10 }} direction='column'>
          <Grid item md={10}>
            <BasicInfoForm />
          </Grid>
          <Grid item md={2}>
            <Box sx={{ display: 'flex', flexDirection: 'row', pt: 2, justifyContent: 'end' }}>
              <Box sx={{ flex: '1 1 auto' }} />

              <Button
                size='large'
                variant="contained"
                color={"inherit"}
                onClick={handleCancel}
                sx={{ mr: 5 }}
              >
                取消
              </Button>
              <Button
                size='large'
                onClick={onSubmit}
                sx={{ mr: 1 }}
                variant="contained"
              >
                更新
              </Button>
            </Box>
          </Grid>

        </Grid>
      </Stack>
    </Grid >
  )
}

export { TunnelUpdateTabs };


