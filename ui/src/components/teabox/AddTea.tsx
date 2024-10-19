import styles from './AddTea.css'
import { useAddTea } from '@/lib/api/teas'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'
import { FaPlus } from 'react-icons/fa'
import { Textarea } from '../common/Textarea'
import { KeyboardEventHandler, useState } from 'react'
import { useGetTeapodInfo } from '@/lib/api/teapods'
import { format } from '@/lib/utility/json'

type Form = {
  data: string
}
const useAddTeaForm = (teapod: string, teabox: string) => {
  const info = useGetTeapodInfo(teapod)
  const addTea = useAddTea(teapod, teabox)
  const form = useForm<Form>()

  const submit = form.handleSubmit(req => addTea.mutate(JSON.parse(req.data)))
  const reset = () => {
    addTea.reset()
    form.reset()
  }

  const error = addTea.data?.error
  const hasError = error !== undefined

  if (info.isSuccess) {
    const placeholder = info.data?.teaboxes.find(b => b.name === teabox)?.placeholder ?? '{}'
    form.setValue('data', format(placeholder))
  }
  if (addTea.isSuccess && !hasError) {
    reset()
  }

  return { ...form, submit, reset, error, hasError }
}

type Props = {
  teapod: string
  teabox: string
}
export const AddTea = ({ teapod, teabox }: Props) => {
  const [open, setOpen] = useState(false)
  const form = useAddTeaForm(teapod, teabox)

  const handleKeyUp: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    e.currentTarget.value = format(e.currentTarget.value)
  }

  return (
    <Dialog.Root open={open || form.hasError} onOpenChange={setOpen}>
      <Dialog.Trigger>
        <Button variant='outline' radius='full' mx='2' style={{ padding: '10px' }}>
          <FaPlus />
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Tea</Dialog.Title>

        {form.hasError && 
          <Callout.Root>
            <Callout.Text>{form.error}</Callout.Text>
          </Callout.Root>
        }

        <form onSubmit={form.submit}>
          <Textarea label='data' onKeyUp={handleKeyUp} className={styles.texts} {...form.register('data')} />

          <div className={styles.actions}>
            <Dialog.Close>
              <Button variant='soft' color='gray' onClick={form.reset}>
                Cancel
              </Button>
            </Dialog.Close>

            <Dialog.Close>
              <Button type='submit'>Save</Button>
            </Dialog.Close>
          </div>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  )
}