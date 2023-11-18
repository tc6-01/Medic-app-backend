import { Tunnel } from 'src/service/models';
import { addTunnel } from 'src/service/tunnelApi';
import { CreateProps } from 'src/types/ComponentProps';
import { TunnelCreateDialogStore, FileSharePageStore } from 'src/states/page/file-share';
import { Box, Button, Grid, Stack } from '@mui/material';
import { useStore } from 'reto';
import { SelectInput, TextInput } from '../src/components/common/InputControls';
import Toast from '../src/components/common/Toast';



const BasicInfoForm = () => {
  const states = useStore(TunnelCreateDialogStore);
  return (
    <Stack spacing={5} p={3} alignItems='space-around'>
      <TextInput name={'名称'} placeHolder={'请输入名称'} onChange={(e) => { states.setTunnelName(e.target.value) }} />
      <TextInput name={'数据通道详情'} placeHolder={'请输入描述'} onChange={(e) => { states.setTunnelDes(e.target.value) }} />
      <SelectInput name={'绑定策略'} options={states.policySelectData} placeHolder={'选择要绑定的策略'} onChange={(e) => { states.setPolicyId(e.target.value) }} />
      {/* <SelectInput name={'拥有者'} options={['拥有者1']} placeHolder={'选择要绑定的拥有者'} onChange={onChange}/> */}
    </Stack>
  )
}

// // const AddAuthorizationForm = () => {
// //     return (
// //         <Stack spacing={5} p={3} alignItems='space-around'>
// //             <TextInput name={'名称'} placeHolder={'请输入名称'} />
// //             <SelectInput name={'绑定策略'} options={['策略1']} placeHolder={'选择要绑定的策略'} />
// //             <SelectInput name={'拥有者'} options={['拥有者1']} placeHolder={'选择要绑定的拥有者'} />
// //         </Stack>
// //     )
// // }

// // const CompleteForm = () => {
// //     return (
// //         <Stack spacing={5} p={3} alignItems='space-around'>
// //             <TextInput name={'名称'} placeHolder={'请输入名称'} />
// //             <SelectInput name={'绑定策略'} options={['策略1']} placeHolder={'选择要绑定的策略'} />
// //             <SelectInput name={'拥有者'} options={['拥有者1']} placeHolder={'选择要绑定的拥有者'} />
// //         </Stack>
// //     )
// // }


// getAllTunnel回调函数 添加数据后会更新数据
const TunnelCreateTabs = ({ getAll }: CreateProps) => {
  const states = useStore(TunnelCreateDialogStore)
  const dialogStates = useStore(FileSharePageStore)

  const totalSteps = () => {
    return steps.length;
  };

  const completedSteps = () => {
    return Object.keys(states.completed).length;
  };

  const isLastStep = () => {
    return states.activeStep === totalSteps() - 1;
  };

  const allStepsCompleted = () => {
    return completedSteps() === totalSteps();
  };

  const handleNext = () => {
    const newActiveStep =
      isLastStep() && !allStepsCompleted()
        ? // It's the last step, but not all steps have been completed,
        // find the first step that has been completed
        steps.findIndex((step, i) => !(i in states.completed))
        : states.activeStep + 1;
    if (isLastStep()) {
      states.setActiveStep(newActiveStep);
      dialogStates.setEditorDialogOpen(false);
    }
    states.setActiveStep(newActiveStep);
  };

  const handleBack = () => {
    states.setActiveStep((prevActiveStep) => prevActiveStep - 1);
  };

  const handleStep = (step: number) => () => {
    states.setActiveStep(step);
  };

  const handleComplete = () => {
    const newCompleted = states.completed;
    newCompleted[states.activeStep] = true;
    states.setCompleted(newCompleted);
    handleNext();
  };

  //取消弹框
  const handleCancel = () => {
    dialogStates.setEditorDialogOpen(false);
  }

  //创建数据通道请求
  const onSubmit = async () => {
    let data: Tunnel = {
      name: states.tunnelName,
      description: states.tunnelDes,
      policyId: states.policyId
    }
    console.log(data)
    let res = await addTunnel(data)

    // if(res?.code == 200)
    // 这个报错返回类型还没解决
    if (res?.status == 0) {
      let msg = res.message;
      dialogStates.setEditorDialogOpen(false);
      getAll();
      Toast.success(msg)
    } else {
      Toast.error('请求失败！')
    }
  }

  const steps = ['基本信息', '添加权限', '完成添加'];
  return (
    <Grid sx={{ minHeight: 500, height: '100%', pl: 10, pr: 10 }} >
      <Stack sx={{ height: 400 }}>
        {/* <Stepper nonLinear activeStep={states.activeStep}>
                    {steps.map((label, index) => (
                        <Step key={label} completed={states.completed[index]}>
                            <StepButton color="inherit" onClick={handleStep(index)}>
                                {label}
                            </StepButton>ß
                        </Step>
                    ))}
                </Stepper>*/}
        <Grid container xs={12} sx={{ pt: 10 }} direction='column'>
          <Grid item md={10}>
            <BasicInfoForm />
          </Grid>
          <Grid item md={2}>
            <Box sx={{ display: 'flex', flexDirection: 'row', pt: 2, justifyContent: 'end' }}>
              <Box sx={{ flex: '1 1 auto' }} />
              {/* {/* <Button
                                size='large'
                                variant="contained"
                                color={states.activeStep === 0 ? "inherit" : "primary"}
                                disabled={states.activeStep === 0}
                                onClick={handleBack}
                                sx={{ mr: 5 }}
                            >
                                上一步
                            </Button>
                            <Button
                                size='large'
                                onClick={handleNext}
                                sx={{ mr: 1 }}
                                variant="contained"
                            >
                                {states.activeStep !== steps.length ? '下一步' : '完成'}
                            </Button> */}
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
                添加
              </Button>
            </Box>
          </Grid>

        </Grid>
      </Stack>
    </Grid >
  )
}

export { TunnelCreateTabs };

