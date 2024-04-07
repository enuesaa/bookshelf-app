import { useDeleteTea } from '@/lib/api/teas'
import { Button, Flex, AlertDialog } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'

type Props = {
  teapod: string
  teaid: string
}
export const DeleteTeaModal = ({ teapod, teaid }: Props) => {
  const deleteTea = useDeleteTea(teapod, teaid)
  const handleDelete: MouseEventHandler<HTMLButtonElement> = (e) => {
    deleteTea.mutate()
  }

  return (
    <AlertDialog.Root>
      <AlertDialog.Trigger>
        <Button color='red' radius='full' style={{ height: '28px' }}>
          delete
        </Button>
      </AlertDialog.Trigger>
      <AlertDialog.Content>
        <AlertDialog.Title>Delete Tea</AlertDialog.Title>
        <AlertDialog.Description size='2'>Do you proceed?</AlertDialog.Description>

        <Flex gap='3' mt='4' justify='end'>
          <AlertDialog.Cancel>
            <Button variant='soft' color='gray'>
              Cancel
            </Button>
          </AlertDialog.Cancel>
          <AlertDialog.Action>
            <Button variant='solid' color='red' onClick={handleDelete}>
              Delete
            </Button>
          </AlertDialog.Action>
        </Flex>
      </AlertDialog.Content>
    </AlertDialog.Root>
  )
}
